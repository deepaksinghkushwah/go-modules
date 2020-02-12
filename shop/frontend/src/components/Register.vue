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
            <input class="form-control" placeholder="Email" type="text" v-model="user.email" />
          </td>
        </tr>
        <tr>
          <td>
            <input class="form-control" placeholder="First Name" type="text" v-model="user.firstname" />
          </td>
        </tr>
        <tr>
          <td>
            <input class="form-control" placeholder="Last Name" type="text" v-model="user.lastname" />
          </td>
        </tr>
        <tr>
          <td>
            <button @click.prevent="register" class="btn btn-primary">Register</button>
          </td>
        </tr>
      </table>
      
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
        password: "",
        email: "",
        firstname: "",
        lastname: "",
      }
    };
  },
  methods: {
    register() {
      this.$store.dispatch("register", this.user);
    },
    
  }
};
</script>