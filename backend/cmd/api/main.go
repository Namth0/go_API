package main

import (
	"log"
	"mon-api/config"
	"mon-api/internal/handlers"
	"mon-api/internal/repository"
	"mon-api/internal/service"
	"mon-api/pkg/middleware"
	"mon-api/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {

	// Utiliser _ pour ignorer dbConfig si nous ne l'utilisons pas
	db, _, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	if db == nil {
		log.Fatal("Failed to initialize database connection")
	}
	log.Println("Successfully connected to database")
	router := gin.Default()
	router.Use(middleware.Logger())

	// Initialisation des dépendances
	articleRepo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepo)
	articleHandler := handlers.NewArticleHandler(articleRepo, articleService)

	// Routes
	v1 := router.Group("/api/v1")
	{
		articles := v1.Group("/articles")
		{
			articles.GET("", articleHandler.GetArticles)
			articles.GET("/:id", articleHandler.GetArticle)
			articles.POST("", articleHandler.CreateArticle)
			articles.PUT("/:id", articleHandler.UpdateArticle)
			articles.DELETE("/:id", articleHandler.DeleteArticle)
		}
	}

	fakeArticles := []models.Article{
		{Title: "Comment Vivaldi est passé de la gloire à l’oubli, puis au statut d'icône", Content: "Les Quatre Saisons sont aussi célèbres qu'un morceau de musique classique peut l'être. Cette collection intemporelle de quatre concertos, chacun incarnant une saison différente, semble aussi entraînante et enivrante aujourd'hui que lorsque Antonio Vivaldi l'a révélée au public en 1725. L'œuvre compte parmi les morceaux de musique classique les plus appréciés jamais composés. Cependant, avant la Seconde Guerre mondiale, seuls quelques spécialistes de l'histoire de la musique en avaient entendu parler. Même le nom de Vivaldi n'était alors qu'une obscure note de bas de page dans certains manuels."},
		{Title: "Voyage : une journée au Caire, l’éclectique capitale égyptienne", Content: "Au Caire, durant la haute saison hivernale, le lever du soleil a lieu un peu avant sept heures. L’accès au site des pyramides de Gizeh, qui se trouve à quinze kilomètres à l’ouest du Caire, ouvre à huit heures toute l’année. Cette possibilité de commencer dès l’aube vous donne davantage de temps et de latitude pour profiter intégralement de la sidération que procurent ces monuments vieux de 4 500 ans et leur compagnon vigilant, le Sphinx de Gizeh. La pyramide de Khéops, seule survivante des Sept merveilles du monde antique, serait constituée de 2,3 millions de blocs de pierre environ."},
		{Title: "Ce que votre âge biologique révèle sur votre santé", Content: "Des chercheurs chinois ont mis au point un outil qui utilise l’intelligence artificielle (IA) pour analyser des images du visage, de la langue et de la rétine afin de déterminer l’âge biologique d’un individu. Cette technologie offre un aperçu de la santé et de l’état de nos cellules, tissus et organes, et de notre prédisposition à développer certaines maladies chroniques."},
	}

	for _, article := range fakeArticles {
		if err := db.Create(&article).Error; err != nil {
			log.Printf("Erreur durant l'insertion des faux articles: %v", err)
		}
	}

	log.Println("Faux articles insérés avec succès!")

	router.Run(":8080")
}
