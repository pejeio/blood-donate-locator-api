const userCredentials = {
    username: "blood",
    password: "donator",
    databaseName: "blood-donate-locator-api",
  };

db.getSiblingDB(userCredentials.databaseName).createUser({
    user: userCredentials.username,
    pwd: userCredentials.password,
    roles: [{ role: "readWrite", db: userCredentials.databaseName }],
});
