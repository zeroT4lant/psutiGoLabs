package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Сервер запущен на порту", port)

	//Отправим сигнал о завершении в консоль
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	//Ждём cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Слушаем до момента пока не получим ctrl+c в консоли
	go func() {
		for {
			select {
			//Если получили сигнал - выходим из цикла
			case <-ctx.Done():
				fmt.Println("Завершение работы слушателя...")
				return
			default:
				//Если не пришёл сигнал, то слушаем-принимаем соединения
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Ошибка при принятии соединения:", err)
					continue
				}

				go handleConnection(ctx, conn)
			}
		}
	}()

	//Когда вышли из цикла, отправляем
	<-stop
	fmt.Println("Получен сигнал завершения работы, закрываем сервер...")
	cancel()

	//Даём время чтобы завершилось выполнение
	time.Sleep(time.Second * 2)
	fmt.Println("Сервер остановлен")
}

func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	fmt.Println("Клиент подключен:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение соединения с:", conn.RemoteAddr())
			return
		default:
			if scanner.Scan() {
				message := scanner.Text()
				fmt.Println("Получено сообщение:", message)

				_, err := conn.Write([]byte("Сообщение получено\n"))
				if err != nil {
					fmt.Println("Ошибка при отправке ответа:", err)
					return
				}
			} else {
				if err := scanner.Err(); err != nil {
					fmt.Println("Ошибка при чтении данных:", err)
				}
				return
			}
		}
	}
}
