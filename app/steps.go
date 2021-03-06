package main

// Request handler for GET, DELETE steps related to specific recipe
// func handleStepsByRecipe(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var steps []Step
// 	var recipe Recipe

// 	db.First(&recipe, params["id"])
// 	db.Model(&recipe).Related(&steps)
// 	recipe.Steps = steps

// 	if r.Method == "GET" {
// 		json.NewEncoder(w).Encode(&recipe)
// 	} else if r.Method == "DELETE" {
// 		db.Delete(&steps)
// 		json.NewEncoder(w).Encode("All steps deleted")
// 	}
// }

// POST new step
// func createStep(w http.ResponseWriter, r *http.Request) {
// 	var steps []Step
// 	var newStep Step

// 	_ = json.NewDecoder(r.Body).Decode(&newStep)

// 	steps = append(steps, newStep)

// 	db.Create(&newStep)

// 	json.NewEncoder(w).Encode(&newStep)
// }

// DELETE single step by stepID
// func deleteStep(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var step Step
// 	var steps []Step
// 	db.First(&step, params["stepID"])
// 	db.Delete(&step)

// 	// returns non-deleted ingredients
// 	db.Find(&steps)
// 	json.NewEncoder(w).Encode(&steps)
// }
