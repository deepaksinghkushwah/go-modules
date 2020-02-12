<template>
  <div id="app">
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <ul class="nav">
        <li class="nav-item"><router-link class="nav-link" to="/">Home</router-link></li>
        <li class="nav-item"><router-link class="nav-link" to="/about">About</router-link></li>
        <li class="nav-item"><router-link class="nav-link" v-if="$store.state.token==null" to="/login">Login</router-link></li>
        <li class="nav-item"><router-link class="nav-link" v-if="$store.state.token==null" to="/register">Register</router-link></li>
        <li class="nav-item"><a class="nav-link" href="javascript:void(0)" v-if="$store.state.token!=null" @click="$store.dispatch('logout')">Logout</a></li>
      </ul>      
    </nav>
    <router-view/>
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
