import {Create, Delete, Get, GetList, GetPersonByAccountIs, Update} from "./common.js";
import {CreateNewAccount} from "./account.js";

const micName = "admins"
const localhostNumber = 8003

export async function GetAdminsList() {
  return await GetList(micName, localhostNumber)
}

export async function GetAdminByAccountID(id) {
  return await GetPersonByAccountIs(id, micName, localhostNumber)
}

export async function GetAdmin(id) {
  return await Get(id, micName, localhostNumber)
}

export async function CreateAdminInfoOnly(name, itn, passport, snils, phone, accountID) {
  const data = {
    name: name.toString(),
    itn: itn.toString(),
    passport: passport.toString(),
    snils: snils.toString(),
    phone: phone.toString(),
    account_id: parseInt(accountID)
  };
  return await Create(data, micName, localhostNumber)
}

export async function UpdateAdmin(id, name, itn, passport, snils, phone, accountID) {
  const data = {
    name: name.toString(),
    itn: itn.toString(),
    passport: passport.toString(),
    snils: snils.toString(),
    phone: phone.toString(),
    account_id: parseInt(accountID)
  };
  return await Update(id, data, micName, localhostNumber)
}

export async function DeleteAdmin(id) {
  return await Delete(id, micName, localhostNumber)
}