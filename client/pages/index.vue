<template>
  <v-card>
    <v-card-title>Login</v-card-title>
    <v-card-text>
      <v-alert type="error" v-if="error">
        {{ error }}
      </v-alert>
      <v-form ref="login_form">
        <v-text-field
          v-model="id"
          label="ID"
          :rules="[required]"
        />
        <v-text-field
          v-model="password"
          label="password"
          type="password"
          :rules="[required]"
        />
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-btn text v-on:click="login">Login</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      id: "",
      password: "",
      required: value => !!value || "必須項目です",
      error: '',
    }
  },
  head() {
    return {
      title: 'ログイン',
    }
  },
  methods: {
    async login() {
      try {
        if (this.$refs.login_form.validate()) {
          await this.$auth.loginWith('local', { data: { ID: this.id, Password: this.password } });
        }
      } catch (error) {
        if (error.response) {
          this.error = error.response.data.Msg
        } else {
          this.error = "予期せぬなエラーが発生しました. 問い合わせてください."
        }
      }
    }
  }
}
</script>
