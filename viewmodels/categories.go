package viewmodels

import(

)

type Categories struct {
	Title string
	Active string
	Categories []Category
}

type Category struct {
	ImageURL string
	Title string
	Description string
	IsOrientRight bool
}

func GetGategories() Categories{
	result := Categories{
		Title: "Lemonade Stand Society - Shop",
		Active: "shop",
	}
	juiceCategory := Category{
		ImageURL: "lemon.png",
		Title: "Juices and mixes",
		Description: "explore our wide assortment of juices explore our wide assortment of juices explore our wide assortment of juices explore our wide assortment of juices explore our wide assortment of juices explore our wide assortment of juices explore our wide assortment of juices explore our wide assortment of juices ",
		IsOrientRight: false,
	}

	supplyCategory := Category{
		ImageURL: "kiwi.png",
		Title: "Cups, straws and other supplies",
		Description: "From paper cups to bio-degradable plastic From paper cups to bio-degradable plastic From paper cups to bio-degradable plastic From paper cups to bio-degradable plastic From paper cups to bio-degradable plastic From paper cups to bio-degradable plastic From paper cups to bio-degradable plastic",
		IsOrientRight: true,
	}

	advertiseCategory := Category{
		ImageURL: "pineapple.png",
		Title: "Signs and Advertising",
		Description: "Sure, you could just wait for people to find your stand you could just wait for people to find your stand you could just wait for people to find your stand you could just wait for people to find your stand you could just wait for people to find your stand you could just wait for people to find your stand",
		IsOrientRight: false,
	}

	result.Categories = []Category{
		juiceCategory,
		supplyCategory,
		advertiseCategory,
	}
	return result
}

