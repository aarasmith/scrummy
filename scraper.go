package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/gocolly/colly"
)

type RecipeIngredient struct {
	group, Name, Amount, Unit, Notes string
}

var recipeIngredients []RecipeIngredient

type RecipeIngredientGroup struct {
	Group string
	Ingredients []RecipeIngredient
}

var recipeIngredientGroups []RecipeIngredientGroup

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) { 
		fmt.Println("Visiting: ", r.URL) 
	}) 
	 
	c.OnError(func(_ *colly.Response, err error) { 
		log.Println("Something went wrong: ", err) 
	}) 
	 
	c.OnResponse(func(r *colly.Response) { 
		fmt.Println("Page visited: ", r.Request.URL) 
	}) 

	c.OnHTML("div.wprm-recipe-ingredient-group", func(e *colly.HTMLElement) { 

		recipeIngredientGroup := RecipeIngredientGroup{}
		recipeIngredientGroup.Group = e.ChildText(".wprm-recipe-ingredient-group-name")
		// fmt.Printf("%s\n", recipeIngredientGroup) 
		e.ForEach("li.wprm-recipe-ingredient", func(_ int, el *colly.HTMLElement) {
			recipeIngredient := RecipeIngredient{}
			
			recipeIngredient.group = recipeIngredientGroup.Group
			recipeIngredient.Name = el.ChildText(".wprm-recipe-ingredient-name")
			recipeIngredient.Amount = el.ChildText(".wprm-recipe-ingredient-amount")
			recipeIngredient.Unit = el.ChildText(".wprm-recipe-ingredient-unit")
			recipeIngredient.Notes = el.ChildText(".wprm-recipe-ingredient-notes")

			recipeIngredients = append(recipeIngredients, recipeIngredient)
			
			recipeIngredientGroup.Ingredients = append(recipeIngredientGroup.Ingredients, recipeIngredient)

			// fmt.Printf("%+v\n", recipeIngredientGroup) 
		})

		recipeIngredientGroups = append(recipeIngredientGroups, recipeIngredientGroup)
		// fmt.Printf("%+v\n", recipeIngredientGroup)
		
		empJSON, err := json.MarshalIndent(recipeIngredientGroup, "", "  ")
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("MarshalIndent funnction output %s\n", string(empJSON))
	}) 
	 
	// c.OnHTML("li.wprm-recipe-ingredient", func(e *colly.HTMLElement) { 
	// 	// printing all URLs associated with the a links in the page
	// 	recipeIngredient := RecipeIngredient{}
		
	// 	recipeIngredient.name = e.ChildText(".wprm-recipe-ingredient-name")
	// 	recipeIngredient.amount = e.ChildText(".wprm-recipe-ingredient-amount")
	// 	recipeIngredient.unit = e.ChildText(".wprm-recipe-ingredient-unit")
	// 	recipeIngredient.notes = e.ChildText(".wprm-recipe-ingredient-notes")

	// 	recipeIngredients = append(recipeIngredients, recipeIngredient)

	// 	fmt.Printf("%v\n", recipeIngredient) 
	// }) 
	 
	c.OnScraped(func(r *colly.Response) { 
		fmt.Println(r.Request.URL, " scraped!") 
	})
	

	c.Visit("https://www.recipetineats.com/biryani/")

	fmt.Println("Hello, World!")
}

//wprm-recipe-ingredient-group
//wprm-recipe-ingredient-group-name
//https://www.recipetineats.com/vindaloo/
//https://www.recipetineats.com/biryani/