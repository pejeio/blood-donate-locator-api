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

<br>

<p align="center">
<a href="https://www.buymeacoffee.com/pejeio" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 40px !important;width: 145px !important;" ></a>
</p>
