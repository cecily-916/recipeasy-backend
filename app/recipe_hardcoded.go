package main

var sampleUser = User{
	Name:     "Cecily",
	Username: "Me!",
	Email:    "email@email.com",
	OwnRecipes: []Recipe{
		{Title: "Oatmeal",
			Description: "This is a delicious breakfast",
			Rating:      1,
			PrepTime:    "10 mins",
			Collections: []*Collection{
				{Name: "Breakfast"},
			},
			CookTime:  "15 mins",
			MainImage: "https://www.veggieinspired.com/wp-content/uploads/2015/05/healthy-oatmeal-berries-featured.jpg",
			Steps: []Step{
				{StepOrder: 1,
					Details: "Add oats to the bowl",
					Ingredients: []Ingredient{
						{Amount: 3, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
						{Amount: 10, Ingredient: "salt", UnitMeasurement: "teaspoon", IngredientOrder: 2},
					},
				},
				{StepOrder: 2,
					Details: "This is another step",
					Ingredients: []Ingredient{
						{Amount: 1, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
					},
				},
			}},
		{Title: "Tea",
			Description: "I love tea!",
			Rating:      5,
			PrepTime:    "2 mins",
			CookTime:    "4 mins",
			MainImage:   "https://images.pexels.com/photos/1417945/pexels-photo-1417945.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
			Steps: []Step{
				{
					StepOrder: 1,
					Details:   "Heat up your water",
					Ingredients: []Ingredient{
						{Amount: 3, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
						{Amount: 10, Ingredient: "salt", UnitMeasurement: "teaspoon", IngredientOrder: 2},
					},
				},
				{
					StepOrder: 2,
					Details:   "Add tea bag to the water",
					Ingredients: []Ingredient{
						{Amount: 25, Ingredient: "sugar", UnitMeasurement: "pound", IngredientOrder: 1},
					},
				},
				{
					StepOrder: 3,
					Details:   "Let steep for 3 mins, then serve!",
					Ingredients: []Ingredient{
						{Amount: 25, Ingredient: "sugar", UnitMeasurement: "pound", IngredientOrder: 1},
					},
				},
			},
		},
		{Title: "Thai Oh My Salmon",
			Description: "Salmon over broccoli",
			Rating:      5,
			PrepTime:    "10 mins",
			CookTime:    "20 mins",
			Servings:    2,
			MainImage:   "https://d2gtpjxvvd720b.cloudfront.net/system/recipe/image/2123/retina_hungry-girl-healthy-thai-oh-my-recipe2-20190528-1730-4972-1885.jpg",
			Steps: []Step{
				{
					StepOrder: 1,
					Details:   "Preheat oven to 375 degrees. Lay a large piece of heavy-duty foil on a baking sheet, and spray with nonstick spray.",
				},
				{
					StepOrder: 2,
					Details:   "In a small bowl, combine chili sauce, soy sauce, and 2 tsp. water. Mix until uniform.",
					Ingredients: []Ingredient{
						{Amount: 1, Ingredient: "soy sauce", UnitMeasurement: "tbsp", IngredientOrder: 1},
						{Amount: 3, Ingredient: "sweet Asian chili sauce", UnitMeasurement: "tbsp", IngredientOrder: 2},
						{Amount: 2, Ingredient: "water", UnitMeasurement: "tsp", IngredientOrder: 3},
					},
				},
				{
					StepOrder: 3,
					Details:   "Distribute broccoli onto the center of the foil. Top with salmon, and sprinkle with 1/4 tsp. garlic powder and 1/8 tsp. paprika. Drizzle with sauce mixture",
					Ingredients: []Ingredient{
						{Amount: 2.5, Ingredient: "broccoli florets", UnitMeasurement: "cups", IngredientOrder: 1},
						{Amount: 8, Ingredient: "raw skinless salmon fillets (2)", UnitMeasurement: "oz", IngredientOrder: 2},
						{Amount: .25, Ingredient: "garlic powder", UnitMeasurement: "tsp", IngredientOrder: 3},
						{Amount: .125, Ingredient: "paprika", UnitMeasurement: "tsp", IngredientOrder: 4},
					},
				},
				{
					StepOrder: 4,
					Details:   "Cover with another large piece of foil. Fold together and seal all four edges of the foil pieces, forming a well-sealed packet.",
				},
				{
					StepOrder: 5,
					Details:   "Bake for 20 minutes, or until salmon is cooked through and veggies are tender.",
				},
				{
					StepOrder: 6,
					Details:   "Cut packet to release hot steam before opening entirely",
				},
			},
		},
		{Title: "Ultimate Chocolate Chip Cookies",
			Description:     "We named this recipe “Ultimate Chocolate Chip Cookies,” because it’s got everything a cookie connoisseur could possibly ask for. With a texture that is slightly crispy on the outside and chewy on the inside, it’s a favorite chocolate chip cookie recipe that’s been top-rated by hundreds of satisfied home cooks.",
			Rating:          5,
			OriginalCreator: "Betty Crocker",
			URL:             "https://www.bettycrocker.com/recipes/ultimate-chocolate-chip-cookies/77c14e03-d8b0-4844-846d-f19304f61c57",
			PrepTime:        "15 mins",
			CookTime:        "1 hr 15 mins",
			Servings:        48,
			MainImage:       "https://joyfoodsunshine.com/wp-content/uploads/2018/02/best-chocolate-chip-cookies-recipe-1.jpg",
			Steps: []Step{
				{
					StepOrder: 1,
					Details:   "Heat oven to 375°F. In small bowl, mix flour, baking soda and salt; set aside.",
					Ingredients: []Ingredient{
						{Amount: 2.25, Ingredient: "all-purpose flour", UnitMeasurement: "cup", IngredientOrder: 1},
						{Amount: 1, Ingredient: "baking soda", UnitMeasurement: "tsp", IngredientOrder: 2},
						{Amount: .5, Ingredient: "salt", UnitMeasurement: "tsp", IngredientOrder: 3},
					},
					Image: "https://images-gmi-pmc.edge-generalmills.com/ae7633e8-25b4-403b-a86c-03b7300982dc.jpg",
				},
				{
					StepOrder: 2,
					Details:   "In large bowl, beat softened butter and sugars with electric mixer on medium speed, or mix with spoon about 1 minute or until fluffy, scraping side of bowl occasionally.",
					Ingredients: []Ingredient{
						{Amount: 1, Ingredient: "softened butter", UnitMeasurement: "cup", IngredientOrder: 1},
						{Amount: .75, Ingredient: "granulated sugar", UnitMeasurement: "cup", IngredientOrder: 2},
						{Amount: .75, Ingredient: "packed brown sugar", UnitMeasurement: "cup", IngredientOrder: 3},
					},
					Image: "https://images-gmi-pmc.edge-generalmills.com/2a409724-c7a1-4110-8048-8f592ed65a3b.jpg",
				},
				{
					StepOrder: 3,
					Details:   "Beat in egg and vanilla until smooth. Stir in flour mixture just until blended (dough will be stiff). Stir in chocolate chips and nuts.",
					Ingredients: []Ingredient{
						{Amount: 1, Ingredient: "egg", UnitMeasurement: "", IngredientOrder: 1},
						{Amount: 1, Ingredient: "vanilla", UnitMeasurement: "tsp", IngredientOrder: 2},
						{Ingredient: "flour mixture", UnitMeasurement: "", IngredientOrder: 3},
						{Amount: .125, Ingredient: "paprika", UnitMeasurement: "tsp", IngredientOrder: 4},
					},
					Image: "https://images-gmi-pmc.edge-generalmills.com/466a6689-1547-4148-98ef-44a2b8c318c8.jpg",
				},
				{
					StepOrder: 4,
					Details:   "Onto ungreased cookie sheets, drop dough by rounded tablespoonfuls 2 inches apart.",
					Image:     "https://images-gmi-pmc.edge-generalmills.com/205de745-2319-4cf6-8322-e6fb7d2293a7.jpg",
				},
				{
					StepOrder: 5,
					Image:     "https://images-gmi-pmc.edge-generalmills.com/cbe8b51a-c3c1-4dcf-8d79-76f98565d3e0.jpg",
					Details:   "Bake 8 to 10 minutes or until light brown (centers will be soft). Cool 2 minutes; remove from cookie sheet to cooling rack. Cool completely, about 30 minutes. Store covered in airtight container.",
				},
			},
		},
	},
}
