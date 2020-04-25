<template>
  <v-layout
    column
    justify-center
    align-center
  >
    <v-flex
      xs12
      sm8
      md6
    >
      <v-card>
        <v-card-title>ユーザーページ</v-card-title>
        <v-card-text>
          <button @click="logout">Logout</button>
        </v-card-text>
      </v-card>
      <v-card>
        <v-card-title>アプリケーション管理</v-card-title>
        <v-card-text>
          csgの認証システムを利用したアプリケーションの作成ができます
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  middleware: 'auth',
  created() {
    this.$axios.$post('/api/get_token_authority')
      .then((result) => {
        if (result.Admin) {
          this.$router.push('/admin')
        }
      })
      .catch((error) => {
        this.$router.push('/')
      })
  },
  methods: {
    /**
     * 管理画面からログアウトする(token削除)
     */
    logout() {
      this.$auth.logout()
    },
  }
}
</script>