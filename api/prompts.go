package api

import (
	"math/rand"
)

type Prompts struct {
	DentalTopics struct {
		CrownAndBridges                        []string `json:"crown_and_bridges"`
		CommonOralAnomalies                    []string `json:"common_oral_anomalies"`
		OralCancers                            []string `json:"oral_cancers"`
		OralLesions                            []string `json:"oral_lesions"`
		SystemicDiseasesWithOralManifestations []string `json:"systemic_diseases_with_oral_manifestations"`
		FailureOfRCT                           []string `json:"failure_of_rct"`
	} `json:"dental_topics"`
	PersonalGrowthAndMindfulness struct {
		PersonalGrowth        []string `json:"personal_growth"`
		Mindfulness           []string `json:"mindfulness"`
		AngerManagementTricks []string `json:"anger_management_tricks"`
		PowerOfHabits         []string `json:"power_of_habits"`
	} `json:"personal_growth_and_mindfulness"`
	LifestyleAndHobbies struct {
		BookSuggestionsWithBriefDescription []string `json:"book_suggestions_with_brief_description"`
		MeaningfulMovies                    []string `json:"meaningful_movies"`
		MasteringSkills                     []string `json:"mastering_skills"`
		FoodRecipes                         []string `json:"food_recipes"`
	} `json:"lifestyle_and_hobbies"`
}

var prompts = Prompts{
	DentalTopics: struct {
		CrownAndBridges                        []string `json:"crown_and_bridges"`
		CommonOralAnomalies                    []string `json:"common_oral_anomalies"`
		OralCancers                            []string `json:"oral_cancers"`
		OralLesions                            []string `json:"oral_lesions"`
		SystemicDiseasesWithOralManifestations []string `json:"systemic_diseases_with_oral_manifestations"`
		FailureOfRCT                           []string `json:"failure_of_rct"`
	}{
		CrownAndBridges: []string{
			"Explain the process of getting dental crowns and bridges, highlighting their importance in restoring tooth function and aesthetics.",
			"What are the latest advancements in dental crowns and bridges, and how do they impact treatment outcomes?",
			"Discuss the common complications that may arise with dental crowns and bridges and how they are managed.",
		},
		CommonOralAnomalies: []string{
			"Describe the most frequently encountered oral anomalies in dental practice and their implications for patient care.",
			"What are the genetic and environmental factors that contribute to common oral anomalies?",
			"Provide an overview of how common oral anomalies can affect overall oral function and aesthetics.",
		},
		OralCancers: []string{
			"Outline the risk factors for developing oral cancer and the importance of early detection.",
			"What are the current treatment modalities for oral cancer and their success rates?",
			"Discuss the common signs and symptoms of oral cancer that dental professionals should look out for during routine exams.",
		},
		OralLesions: []string{
			"What are the most commonly diagnosed oral lesions, and how are they typically treated?",
			"Explain the difference between benign and malignant oral lesions, including examples of each.",
			"How can early detection of oral lesions improve patient outcomes?",
		},
		SystemicDiseasesWithOralManifestations: []string{
			"Discuss how systemic diseases like diabetes or cardiovascular conditions manifest in the oral cavity.",
			"What role do dental professionals play in identifying early signs of systemic diseases through oral examinations?",
			"Provide examples of systemic diseases that are often first detected through oral symptoms.",
		},
		FailureOfRCT: []string{
			"What are the common causes of root canal therapy (RCT) failure, and how can they be prevented?",
			"Discuss the options for retreatment or alternatives after a failed RCT.",
			"What are the signs that a root canal treatment is failing, and how should a dentist proceed?",
		},
	},
	PersonalGrowthAndMindfulness: struct {
		PersonalGrowth        []string `json:"personal_growth"`
		Mindfulness           []string `json:"mindfulness"`
		AngerManagementTricks []string `json:"anger_management_tricks"`
		PowerOfHabits         []string `json:"power_of_habits"`
	}{
		PersonalGrowth: []string{
			"What are the key habits that contribute to personal growth and self-improvement?",
			"Discuss the role of self-reflection in personal growth and ways to incorporate it into daily life.",
			"How does setting specific goals aid in personal development and maintaining motivation?",
		},
		Mindfulness: []string{
			"How can mindfulness practices improve mental health and daily productivity?",
			"What are some simple mindfulness exercises to incorporate into a busy daily routine?",
			"Explain the science behind mindfulness and how it affects the brain.",
		},
		AngerManagementTricks: []string{
			"What are effective techniques for managing anger in stressful situations?",
			"How does mindfulness help in controlling anger and emotional responses?",
			"Provide a step-by-step guide for using deep breathing to manage moments of anger.",
		},
		PowerOfHabits: []string{
			"Explain how small daily habits can lead to significant personal and professional growth.",
			"What are the key elements that make a habit stick, and how can we use them to build new habits?",
			"Discuss the role of 'habit stacking' in creating a more productive lifestyle.",
		},
	},
	LifestyleAndHobbies: struct {
		BookSuggestionsWithBriefDescription []string `json:"book_suggestions_with_brief_description"`
		MeaningfulMovies                    []string `json:"meaningful_movies"`
		MasteringSkills                     []string `json:"mastering_skills"`
		FoodRecipes                         []string `json:"food_recipes"`
	}{
		BookSuggestionsWithBriefDescription: []string{
			"Recommend a meaningful book about personal development with a brief description of its key insights.",
			"Suggest a book that explores the theme of habits and how they shape our lives.",
			"What book would you recommend for mastering mindfulness, and why is it impactful?",
		},
		MeaningfulMovies: []string{
			"Recommend a meaningful movie that explores themes of personal growth or overcoming adversity.",
			"What movie would you suggest that promotes mindfulness or the power of self-reflection?",
			"Share a brief description of a film that focuses on mastering one's skills or achieving greatness.",
		},
		MasteringSkills: []string{
			"What are the steps to mastering a new skill, and how can one stay motivated throughout the process?",
			"How does the '10,000-hour rule' apply to mastering complex skills, and what are some examples of this in practice?",
			"Discuss the importance of deliberate practice in skill mastery and provide an example.",
		},
		FoodRecipes: []string{
			"Share a healthy, easy-to-make recipe that boosts energy and promotes well-being.",
			"What are some quick and delicious meals that support mindfulness and focus?",
			"Provide a recipe for a traditional Iraqi dish with a modern twist.",
		},
	},
}

func GetRandomQuestion() string {
	allPrompts := [][]string{
		prompts.DentalTopics.CrownAndBridges,
		prompts.DentalTopics.CommonOralAnomalies,
		prompts.DentalTopics.OralCancers,
		prompts.DentalTopics.OralLesions,
		prompts.DentalTopics.SystemicDiseasesWithOralManifestations,
		prompts.DentalTopics.FailureOfRCT,
		prompts.PersonalGrowthAndMindfulness.PersonalGrowth,
		prompts.PersonalGrowthAndMindfulness.Mindfulness,
		prompts.PersonalGrowthAndMindfulness.AngerManagementTricks,
		prompts.PersonalGrowthAndMindfulness.PowerOfHabits,
		prompts.LifestyleAndHobbies.BookSuggestionsWithBriefDescription,
		prompts.LifestyleAndHobbies.MeaningfulMovies,
		prompts.LifestyleAndHobbies.MasteringSkills,
		prompts.LifestyleAndHobbies.FoodRecipes,
	}

	categoryIndex := rand.Intn(len(allPrompts))
	selectedCategory := allPrompts[categoryIndex]
	questionIndex := rand.Intn(len(selectedCategory))
	return selectedCategory[questionIndex]
}
