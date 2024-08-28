# Groupie Tracker Filters

Groupie Tracker is a web application that allows users to explore information about various musical artists, including their members, creation date, first album release, locations, concert dates, and relationships between artists.

## Features

- **Artist List**: View a list of all the artists in the database, including their name, image, and basic details.
- **Artist Details**: Click on an artist to see more detailed information about them, including their members, creation date, first album release, locations, concert dates, and relationship to other artists.
- **Filtering**: Filter the list of artists by various criteria, such as name, creation date, or location.
- **Search**: Search for artists by name or other details.
- **Geolocation**: View the locations of an artist's concerts on a map.

## Project Structure

The project is organized into the following directories:

- [datatypes/](datatypes/): Contains Go structs used to represent the various types of data.
- [handlers/](handlers/): Contains the Go handlers that handle HTTP requests and responses.
- [templates/](templates/): Contains the HTML templates used to render the user interface.
- [templates/scripts/](templates/scripts/): Contains the JavaScript files used to add interactivity to the user interface.
- [utils/](utils/): Contains various utility functions used throughout the application.
- [json/](json/): Contains the JSON files that store the artist, location, and concert date data.
- [main.go](main.go): The main program of the application.

## Getting Started

To run the Groupie Tracker application, follow these steps:

1. Ensure you have Go installed on your system.
2. Clone the repository to your local machine by running the following command in your terminal:
   ```bash
   git clone https://zone01normandie.org/git/faoudia/groupie-tracker.git
   ```
3. Navigate to the project directory in your terminal.
4. Run the application using the 
    ```bash 
    go run main.go
    ```
5. Open your web browser and navigate to [http://localhost:8080](http://localhost:8080) to access the application.