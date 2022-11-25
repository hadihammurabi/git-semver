package service

import "github.com/hadihammurabi/go-ioc/ioc"

func Build() {
	ioc.Set(NewSemverService)
	ioc.Set(NewGitService)
}
