package main

import (
	"go.uber.org/fx"
	"log"
)

// ==================================Main Service==============================================

type MainService struct {
	publisher iPublisher
}

func NewMainService(publisher iPublisher) *MainService {
	return &MainService{publisher: publisher}
}

func (service *MainService) Run() {
	service.publisher.Publish()
	log.Print("The Main Service started")
}

// ===================================Dependency===============================================
type iPublisher interface {
	Publish()
}

type Publisher struct {
	titles []*Title
}

func NewPublisher(titles ...*Title) *Publisher {
	return &Publisher{titles: titles}
}

func (publisher *Publisher) Publish() {
	for _, title := range publisher.titles {
		log.Print("Publisher ", *title)
	}

}

// ===============================Dependency on publisher======================================

type Title string

func main() {
	//t := Title(" hello")
	//np := NewPublisher(&t)
	//ms := NewMainService(np)
	//ms.Run()

	fx.New(
		fx.Provide(NewMainService),
		fx.Provide(
			fx.Annotate(
				NewPublisher,
				fx.As(new(iPublisher)),
				fx.ParamTags(`group:"titles"`),
			),
			fx.Provide(
				titleComponents("title one")),
			fx.Provide(
				titleComponents("title two"),
			),
			fx.Invoke(func(service *MainService) {
				service.Run()
			}))).Run()

}

func titleComponents(title any) any {
	return fx.Annotate(
		func() *Title {
			title := Title("title two")
			return &title
		},
		fx.ResultTags(`group:"titles"`))
}
