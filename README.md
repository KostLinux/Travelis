# Travelis

Travelis is a web application that enables users to book their journeys with ease and access the offers they need in just one click.

The application automatically calculates the offers, which can be sorted through multiple steps. Additionally, the offers include the calculation of Revolut cashback (with the option to disable it).

## Requirements

- [Revolut account](https://www.revolut.com)
- [Google account](https://accounts.google.com/)
- [Git](https://git-scm.com/downloads)
- [Code Editor](https://code.visualstudio.com)
- [GoLang](https://go.dev)
- [Docker](https://docs.docker.com/get-docker/)
- [KeepassXC](https://keepassxc.org)

# API

API's used for web application:

- Air flight API: [SkyScanner](https://developers.skyscanner.net)
- Restaurants API: [Yelp](https://docs.developer.yelp.com/docs/getting-started)
- Hotels Booking API: [Booking](https://developers.booking.com/landing/)
- Revolut Booking API with cashback (availability to be confirmed)

## Objectives

Create a web application where users can input the following information:

- Departure date
- Arrival date
- Number of persons
- Departure city / country
- Arrival city / country
- Fuel consumption (optional)

The application will parse information from multiple APIs and provide offers that include the following in a consolidated manner:

- Air flight prices
- Hotel prices
- Public transport prices
- Best-rated restaurants according to user ratings

Additionally, the application will calculate the daily money needs, which will include:

- Restaurant costs per day (estimated at 50-75€)
- Average souvenir price (estimated at 30-50€)

The application should offer the following options:

- Sorting: Users can sort offers by descending/ascending prices, dates, and star ratings.
- Include/Exclude: Users can choose to include or exclude specific services, such as:
    - air flights
    - automated money needs calculation 
    - Revolut card cashback
    - and public transport
    - hotels
    - car fuel consumption (Family version only)

When excluding air flights and hotels, the application will primarily focus on calculating costs related to restaurants and souvenirs. This feature allows users to estimate their expenses specifically for dining out and purchasing souvenirs during their travels, providing them with a clearer understanding of their expected costs and enabling them to budget accordingly. By narrowing the scope to these categories, the application aims to offer a more tailored and personalized experience, empowering travelers to make informed decisions about their spending.

It is important to note that while air flights and hotels are not included in the cost calculation, it is still highly recommended to consider these aspects when planning a trip to ensure comprehensive budgeting and a seamless travel experience. These components play significant roles in the overall journey and should be taken into account for a well-rounded travel plan.

Moreover, the application incorporates fuel consumption calculations for users who plan to travel by car. This feature allows travelers to estimate the fuel costs for their entire trip, considering the distance between Country 1 and Country 2, as well as the return journey from Country 2 to Country 1. This functionality adds convenience and helps users make informed decisions when deciding on their mode of transportation.
## Version

Application will have 2 versions:

- Free version

    - Users have access to offers; however, these offers will be sent via email with links to other companies, such as flight and hotel providers.    

- Family version (15$ per month)
    - Users have access to offers, allowing them to conveniently book planes and hotels (if they are not excluded).
    - Users can benefit from the fuel consumption calculation feature for car travel.

These versions provide users with various options based on their preferences and requirements, ensuring a customized experience that suits their individual needs.

Please note that all prices and features mentioned are subject to change and may vary based on the application's specific offerings and business model.

# Contributing

You can find the detailed contributing guidelines by navigating to the [`CONTRIBUTE.md`](/CONTRIBUTE.md/) file in the repository.

# Trello

You can look for tasks in [Trello](https://trello.com/b/rfrdh8iP/travelis-kanban)

# Docs

- [Local-Env](/doc/local-env.md/) - Setting up Local Environment