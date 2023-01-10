<h1 align="center">🩸<br>Blood Donate Locator API</h1>
<p align="center">
	<b>Open source API</b> with worldwide locations where donors are welcome to <b>donate blood</b> or blood plasma.
</p>
<br>

## 👀 Prerequisites
- Make sure you have a [MongoDB](https://www.mongodb.com/) database
- Make sure you have Go installed ([download](https://go.dev/dl/)). Version `1.14` or higher is required.
## ⚙️ Installation


1. Install dependencies
	```sh
	go mod tidy
	```
2. Create the `.env` file

	```sh
	# Edit the file to your needs
	cp .env.sample .env
	```
3. Create the `auth-users.csv` file

	Authentication is implemented to validate if a user can update resources.
	```sh
	# Edit the file to your needs
	cp auth-users.csv.sample auth-users.csv
	```
4. Create the `casbin_policy.csv` file

	Authorization is implemented to validate if a user can update resources.
	```sh
	# Edit the file to your needs
	cp casbin_policy.csv.sample casbin_policy.csv
	```
5. Run the project with [Air](https://github.com/cosmtrek/air) *(live reload for Go apps)*


## 🔨 Built With

* [Fiber](https://github.com/gofiber/fiber)
* [Go](https://go.dev)
* [Logrus](https://github.com/sirupsen/logrus)
* [Viper](https://github.com/spf13/viper)

## ⚠️ License
1. [LICENSE](LICENSE)


## 👷 Roadmap
- App with map view: This would allow users to see the locations of blood donation centers on a map, making it easier for them to find a location that is convenient for them.
- Search functionality: Users should be able to search for blood donation centers by location, such as by city or zip code.
- Filtering options: Users should be able to filter the list of blood donation centers by various criteria, such as by type of donation (whole blood, platelets, etc.), by location (e.g., within a certain distance of their current location), or by other factors such as age or weight requirements.
- Information about each location: The app should provide detailed information about each blood donation center, including its address, hours of operation, and any special requirements or restrictions.
- Appointments: Users should be able to schedule appointments to donate blood at a particular location, either through the app or by contacting the center directly.
- Notifications: The app could send users reminders about upcoming blood drives or other events, or alert them when there is an urgent need for blood donations in their area.

<br>

<p align="center">
<a href="https://www.buymeacoffee.com/pejeio" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 40px !important;width: 145px !important;" ></a>
</p>
