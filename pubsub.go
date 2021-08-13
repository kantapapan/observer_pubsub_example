package main

import (
	"container/list"
	"fmt"
)

// Observer ...
type Observer interface {
	Notify(data interface{})
}

// Observable ...
type Observable struct {
	subs *list.List
}

// Subscribe 購読処理
func (o *Observable) Subscribe(observer Observer) {
	o.subs.PushBack(observer)
}

// Unsubscribe 購読停止
func (o *Observable) Unsubscribe(observer Observer) {
	for sub := o.subs.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Observer) == observer {
			o.subs.Remove(sub)
		}
	}
}

// Fire 発火処理
func (o *Observable) Fire(data interface{}) {
	for sub := o.subs.Front(); sub != nil; sub = sub.Next() {
		sub.Value.(Observer).Notify(data)
	}
}

///===

// Publisher 出版社
type Publisher struct {
	Observable
	Title string
}

// NewPublisher ...
func NewPublisher(title string) *Publisher {
	return &Publisher{
		Observable: Observable{subs: new(list.List)},
		Title:      title,
	}
}

// Topic トピック
func (p *Publisher) Topic() {
	p.Fire(p.Title)
}

///===

// GameService ゲーム
type GameService struct{}

// Notify ...
func (b *GameService) Notify(data interface{}) {
	fmt.Printf("Books content : %s\n", data.(string))
}

// BlogService 本
type BlogService struct{}

// Notify ...
func (b *BlogService) Notify(data interface{}) {
	fmt.Printf("Blog content : %s\n", data.(string))
}

// MovieService 映画
type MovieService struct{}

// Notify ...
func (b *MovieService) Notify(data interface{}) {
	fmt.Printf("Movie content : %s\n", data.(string))
}

func main() {

	// 出版社
	p := NewPublisher("ELDEN RING 2022年1月21日に発売")

	gameService := &GameService{}
	blogService := &BlogService{}
	MovieService := &MovieService{}

	p.Subscribe(gameService)
	p.Subscribe(blogService)
	p.Subscribe(MovieService)

	p.Topic()

	fmt.Println("---")

	p.Unsubscribe(gameService)

	p.Topic()

}
