<script setup>
import HelloWorld from './components/HelloWorld.vue'
// ----------------- ACCOUNT -------------------

// import {
//   CreateNewAccount,
//   DeleteAccount,
//   GetAccountByID,
//   GetAccountByLoginAndPassword,
//   UpdateAccount
// } from "./api/account.js";

// вход - возвращает роль и id аккаунта
// на основе роли мы принимаем решение о том, к какому микросервису обращаться с полученным id
// не знаю, так ли это работает на самом деле, но чат гпт сказал, что так
// console.log(await GetAccountByLoginAndPassword("admin1", "adminpass1"))

// console.log(await GetAccountByID(1))

// названия ролей: admin - админ | employee - сотрудник | owner - владелец точки
// console.log(await CreateNewAccount("new", "1111", "admin"))

// console.log(await UpdateAccount(6, "", "", ""))

// console.log(await DeleteAccount(6))

// ----------------- ADMIN -------------------

// console.log(await GetAdmin(1))
// console.log(await GetAdminsList())

</script>

<template>
  <div>
    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div>
  <HelloWorld msg="Vite + Vue" />
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
