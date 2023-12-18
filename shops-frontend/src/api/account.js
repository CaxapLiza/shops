import {Create, Delete, Get, LogIn, Update} from "./common.js";

const micName = "accounts";
const localhostNumber = 8000;

export async function GetAccountByLoginAndPassword(login, password) {
  await LogIn(login, password, micName, localhostNumber);
}

export async function GetAccountByID(id) {
  return await Get(id, micName, localhostNumber);
}

export async function CreateNewAccount(login, password, role) {
  const data = {
    login: login.toString(),
    password: password.toString(),
    role: role.toString()
  };
  return await Create(data, micName, localhostNumber);
}

export async function UpdateAccount(id, login, password, role) {
  const data = {
    login: login.toString(),
    password: password.toString(),
    role: role.toString()
  };
  return await Update(id, data, micName, localhostNumber);
}

export async function DeleteAccount(id) {
  return await Delete(id, micName, localhostNumber);
}