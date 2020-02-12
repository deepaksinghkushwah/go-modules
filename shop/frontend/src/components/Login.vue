<template>
  <div class="d-flex align-items-center flex-column justify-content-center h-100">
    <div class="display-4">
      <table class="table table-bordered" v-if="$store.state.token==null">
        <tr>
          <td>
            <input class="form-control" placeholder="Username" type="text" v-model="user.username" />
          </td>
        </tr>
        <tr>
          <td>
            <input
              class="form-control"
              placeholder="Password"
              type="password"
              v-model="user.password"
            />
          </td>
        </tr>
        <tr>
          <td>
            <button @click.prevent="login" class="btn btn-primary">Login</button>
          </td>
        </tr>
      </table>
      <button v-if="$store.state.token!=null" class="btn btn-danger" @click="logout">Logout</button>
    </div>
  </div>
</template>
<script>
export default {
  name: "Login",
  data() {
    return {
      user: {
        username: "",
        password: ""
      }
    };
  },
  methods: {
    login() {
      this.$store.dispatch("login", this.user).then(()=> {
          this.$router.push("/");
      });
    },
    logout(){
        this.$store.dispatch("logout").then(()=> {
          this.$router.push("/");
      });
    }
  }
};
</script>