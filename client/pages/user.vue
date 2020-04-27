<template>
  <v-layout column>
      <v-card>
        <v-card-title>ユーザーページ</v-card-title>
        <v-card-text>
          <button @click="logout">Logout</button>
        </v-card-text>
      </v-card>
      <br>
      <br>
      <v-card>
        <v-card-title>アプリケーション作成</v-card-title>
        <v-card-text>
          <v-alert type="error" v-if="error">
            {{ error }}
          </v-alert>
          <v-form ref="createAppForm">
            <v-text-field
              v-model="id"
              label="アプリケーションID (15文字以下, 英数字)"
              :rules="[required, idLength, alphaNumCheck]"
            />
            <v-text-field
              v-model="name"
              label="アプリケーション名"
              :rules="[required]"
            />
            <v-text-field
              v-model="description"
              label="アプリケーション説明"
              :rules="[required]"
            />
            <v-text-field
              v-model="callback"
              label="コールバックURL"
              :rules="[required]"
            />
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-btn text v-on:click="create">送信する</v-btn>
        </v-card-actions>
      </v-card>
      <br>
      <br>
      <v-card>
        <v-card-title>アプリケーション一覧</v-card-title>
        <v-card-text>
          <v-data-table
            :headers="appsHeaders"
            :items="appsDesserts"
            :items-per-page="5"
            class="elevation-1"
          >
            <template v-slot:item.actions="{ item }">
              <v-icon
                small
                class="mr-2"
                @click="deleteAccount(item)"
              >
                mdi-delete
              </v-icon>
            </template>
          </v-data-table>
        </v-card-text>
      </v-card>
  </v-layout>
</template>

<script>
export default {
  data() {
    return {
      id: "",
      name: "",
      description: "",
      callback: "",
      error: "",
      required: value => !!value || "必須項目です",
      idLength: value => value.length <= 15 || "15文字以内で入力してください",
      alphaNumCheck: value => {
        if (!value.match(/^[A-Za-z0-9]*$/))
          return "半角英数字のみで入力してください"
      },
      // urlCheck: value => {
      //   if (!value.match())
      //     return "URLを入力してください"
      // },
      appsHeaders: [
        {
          text: "ID",
          align: "start",
          value: "ID",
        },
        {
          text: "Name",
          value: "Name",
        },
        {
          text: "Description",
          value: "Description",
        },
        {
          text: "ConsumerKey",
          value: "ConsumerKey",
        },
        {
          text: "ConsumerSecret",
          value: "ConsumerSecret",
        },
        {
          text: "Callback",
          value: "Callback",
        },
      ],
      appsDesserts: [],
    }
  },
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

    this.$axios.$post('/api/user/get_app')
      .then((result) => {
        this.appsDesserts = result
      })
      .catch((error) => {
        console.log(error.response.data.Msg)
      })
  },
  methods: {
    /**
     * 管理画面からログアウトする(token削除)
     */
    logout() {
      this.$auth.logout()
    },
    create() {
      if (this.$refs.createAppForm.validate()) {
        const params = new URLSearchParams()
        params.append('id', this.id)
        params.append('name', this.name)
        params.append('description', this.description)
        params.append('callback', this.callback)
        this.$axios.$post('/api/user/create_app', params)
          .then((result) => {
            console.log(result)
            this.error = ""
            this.id = ""
            this.name = ""
            this.description = ""
            this.callback = ""
            this.$refs.createAppForm.resetValidation()
          })
          .catch((e) => {
            if (e.response) {
              this.error = e.response.data.Msg
            } else {
              this.error = "予期せぬなエラーが発生しました. 問い合わせてください."
            }
          })
      }
    },
  }
}
</script>