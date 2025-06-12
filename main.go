package main

import (
	"html/template"
	"log"
	"net/http"
	"os" // Still needed for PORT environment variable
)

// PageData holds the data to be passed to the HTML template
type PageData struct {
	Version string
	Message string
}

func main() {
	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appVersion := "1.0.0"

	// Define the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Version: appVersion,
			Message: "Hello!",
		}

		// Parse the HTML template
		tmpl, err := template.New("index").Parse(htmlTemplate)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}

		// Execute the template with the data
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	})

	log.Printf("Application Version: %s", appVersion)
	log.Printf("Server listening on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// htmlTemplate is the HTML content embedded in the Go application
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Message}} - Version {{.Version}}</title>
    <style>
        body {
            margin: 0;
            padding: 0;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            background: linear-gradient(135deg, #6dd5ed, #2193b0);
            color: white;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
            overflow: hidden;
        }
        .container {
            text-align: center;
            background: rgba(255, 255, 255, 0.15);
            backdrop-filter: blur(10px);
            border-radius: 20px;
            padding: 50px 70px;
            box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
            border: 1px solid rgba(255, 255, 255, 0.18);
            animation: fadeIn 1.5s ease-out;
        }
        h1 {
            font-size: 6em; /* Extremely big text */
            margin: 0;
            letter-spacing: 5px;
            font-weight: 900;
            position: relative;
            animation: bounceIn 1.2s ease-out;
        }
        h1::before {
            content: '';
            position: absolute;
            top: -20px;
            left: -20px;
            right: -20px;
            bottom: -20px;
            background: rgba(255, 255, 255, 0.1);
            border-radius: 15px;
            z-index: -1;
            filter: blur(15px);
        }
        p {
            font-size: 2.5em;
            margin-top: 20px;
            font-weight: 300;
            animation: slideInUp 1.5s ease-out;
        }

        @keyframes fadeIn {
            from { opacity: 0; transform: translateY(20px); }
            to { opacity: 1; transform: translateY(0); }
        }

        @keyframes bounceIn {
            0%, 20%, 40%, 60%, 80%, 100% {
                animation-timing-function: cubic-bezier(0.215, 0.610, 0.355, 1.000);
            }
            0% {
                opacity: 0;
                transform: scale3d(.3, .3, .3);
            }
            20% {
                transform: scale3d(1.1, 1.1, 1.1);
            }
            40% {
                transform: scale3d(.9, .9, .9);
            }
            60% {
                opacity: 1;
                transform: scale3d(1.03, 1.03, 1.03);
            }
            80% {
                transform: scale3d(.97, .97, .97);
            }
            100% {
                opacity: 1;
                transform: scale3d(1, 1, 1);
            }
        }

        @keyframes slideInUp {
            from { opacity: 0; transform: translateY(50px); }
            to { opacity: 1; transform: translateY(0); }
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>{{.Message}}</h1>
        <p>Application Version: {{.Version}}</p>
    </div>
</body>
</html>
`
