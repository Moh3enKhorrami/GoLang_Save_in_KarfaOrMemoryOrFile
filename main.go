package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	var storeType string
	var filePath string
	var kafkaIP string

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Save to Kafka",
	}

	var storeCmd = &cobra.Command{
		Use:   "store",
		Short: "Data storage in Kafka or File",
		Run: func(cmd *cobra.Command, args []string) {
			var stack StackStorage
			var err error

			if storeType == "file" {
				if filePath == "" {
					fmt.Println("The file address is required.")
					return
				}
				stack, err = NewFile(filePath)
				if err != nil {
					log.Fatalf("Error creating file: %v", err)
				}
				defer func(file *os.File) {
					err := file.Close()
					if err != nil {

					}
				}(stack.(*FileStack).file)

			} else if storeType == "kafka" {
				if kafkaIP == "" {
					fmt.Println("Kafka IP address is required.")
					return
				}
				stack, err = NewKafkaProducer([]string{kafkaIP})
				if err != nil {
					log.Fatalf("Error connecting to Kafka: %v", err)
				}
				defer func(producer sarama.SyncProducer) {
					err := producer.Close()
					if err != nil {

					}
				}(stack.(*KafkaStack).producer)

			} else {
				stack = &MemoryStack{}
			}

			var n int
			fmt.Println("Enter numbers (negative number to stop):")
			for {
				fmt.Print("Enter a number: ")
				fmt.Scanln(&n)
				if n < 0 {
					break
				}
				err := stack.Push(n)
				if err != nil {
					fmt.Println("Error saving: ", err)
				}
			}

			fmt.Println("Popping items from stack:")
			for {
				item, err := stack.Pop()
				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Println(item)
			}
		},
	}

	storeCmd.Flags().StringVarP(&storeType, "store", "s", "", "Storage type: kafka or file")
	storeCmd.Flags().StringVarP(&filePath, "file", "f", "", "File address for storage (if file is selected)")
	storeCmd.Flags().StringVarP(&kafkaIP, "kafka-ip", "k", "", "IP address for Kafka (if kafka is selected)")

	rootCmd.AddCommand(storeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
