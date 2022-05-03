package service

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"task/model"
	services "task/proto"
)

func CreateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskResponse) error {
	ch, err := model.MQ.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	if err != nil {
		return err
	}
	body, _ := json.Marshal(req)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType: "application/json",
		Body: body,
	})
	if err != nil {
		return err
	}
	return nil
}