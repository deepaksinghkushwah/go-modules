import Vue from 'vue'
import Vuex from 'vuex'
import axios from "axios"
import qs from "qs"
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    baseUrl: "http://localhost:8082",
    token: localStorage.getItem('token') || null,

  },
  mutations: {
    LOGIN(state, payload){
      state.token = payload.token;
      localStorage.setItem('token', state.token);
      alert(payload.msg);
    },
    LOGOUT(state){
      state.token = null;      
      localStorage.removeItem("token")
    }
  },
  actions: {
    login(context, user){
      
      axios.post(context.state.baseUrl + '/login', qs.stringify({
        username: user.username,
        password: user.password
      }))
      .then(function (response) {
        if(response.data.status == 0){
          alert(response.data.msg);
        } else {
          context.commit("LOGIN", response.data);
        }        
      })
      .catch(function (error) {
        console.log(error);
      });
      
    },
    logout(context){
      context.commit("LOGOUT");
    },
    register(context, user){
      axios.post(context.state.baseUrl + "/register", qs.stringify(user))
      .then(response => {
        alert(response.data.msg);
      })
      .catch(error => {
        console.log(error);
      })
    }
  },
  modules: {
  }
})
