<template>
  <v-card>
    <v-card-title>{{ result }}</v-card-title>
    <v-card-text>
      {{ resultInfo }}
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      result: "",
      resultInfo: "",
    }
  },
  head() {
    return {
      title: 'メール認証',
    }
  },
  mounted() {
    const params = new URLSearchParams()
    params.append('ID', this.$route.query.id)
    params.append('Code', this.$route.query.code)
    this.$axios.$post('/api/authentication', params)
      .then(() => {
        this.result = "認証完了"
        this.resultInfo = "認証が完了しました. 会費を払い管理者にアカウントをアクティベートしてもらうことで登録が完了となります."
      })
      .catch((err) => {
        this.result = "認証失敗"
        console.log(err.response.data.Msg)
      })
  },
}
</script>