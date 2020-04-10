<template>
  <v-card>
    <v-card-title>Login</v-card-title>
    <v-card-text>
       <v-alert type="success" v-if="success">
        {{ success }}
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
      success: '',
    }
  },
  head() {
    return {
      title: 'ログイン',
    }
  },
  mounted() {
    if (this.$route.query.action) {
      switch (this.$route.query.action) {
        case 'register':
          this.success = '確認メールを送信しました.'
          break
      }
    }
  },
  methods: {
    async login() {
      // if (this.$refs.login_form.validate()) {
      //   console.log('login')
      // }
      try {
        await this.$auth.loginWith('local', { data: { ID: this.id, Password: this.password } });
      } catch (error) {
        console.log('errorだよ')
        console.log(error)
      }
    }
  }
}
</script>
