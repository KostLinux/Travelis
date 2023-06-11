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

## API

### Others
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
- Include/Exclude: Users can choose to include or exclude specific services, such as air flights, automated money needs calculation, Revolut card cashback, and public transport.

After success offer it will send order information to email
# Trello

You can look for tasks in [Trello](https://trello.com/b/rfrdh8iP/travelis-kanban)

# Docs

- [Local-Env](/doc/local-env.md/) - Setting up Local Environment