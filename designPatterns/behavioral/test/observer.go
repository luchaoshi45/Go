package test

import (
	"Go/designPatterns/behavioral/observer"
	"time"
)

/*
	               百晓生
		[丐帮]               [明教]
	    洪七公               张无忌
	    黄蓉					韦一笑
	    乔峰				    金毛狮王
*/

func Observer() {
	h1 := observer.NewHero("洪七公", observer.PGaiBang)
	h2 := observer.NewHero("黄蓉", observer.PGaiBang)
	h3 := observer.NewHero("乔峰", observer.PGaiBang)
	h4 := observer.NewHero("张无忌", observer.PMingJiao)
	h5 := observer.NewHero("韦一笑", observer.PMingJiao)
	h6 := observer.NewHero("金毛狮王", observer.PMingJiao)
	bxs := observer.NewBaiXiaoSheng()

	bxs.AddListener(h1)
	bxs.AddListener(h2)
	bxs.AddListener(h3)
	bxs.AddListener(h4)
	bxs.AddListener(h5)
	bxs.AddListener(h6)

	go h1.HandlerEvent()
	go h2.HandlerEvent()
	go h3.HandlerEvent()
	go h4.HandlerEvent()
	go h5.HandlerEvent()
	go h6.HandlerEvent()

	h2.Fight(h4, bxs)
	h4.Fight(h2, bxs)

	time.Sleep(1000 * time.Second)
}
