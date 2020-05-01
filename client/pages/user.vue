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
        <v-card-title>投票</v-card-title>
        <v-card-text>
          <v-alert type="error" v-if="voteError">
            {{ voteError }}
          </v-alert>
          <v-alert type="success" v-if="voteSuccess">
            {{ voteSuccess }}
          </v-alert>
          <v-form ref="voteForm">
            <v-text-field
              v-model="votedID"
              label="ID"
              :rules="[required]"
            />
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-btn text v-on:click="voteAgree">賛成</v-btn>
          <v-btn text v-on:click="voteDisagree">反対</v-btn>
        </v-card-actions>
      </v-card>
      <br>
      <br>
      <v-card>
        <v-card-title>作成</v-card-title>
        <v-card-text>
          <v-alert type="error" v-if="error">
            {{ error }}
          </v-alert>
          <v-form ref="createVoteForm">
            <v-text-field
              v-model="voteTitle"
              label="title"
              :rules="[required]"
            />
            <v-text-field
              v-model="voteDescription"
              label="description"
              :rules="[required]"
            />
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-btn text v-on:click="createVote">送信する</v-btn>
        </v-card-actions>
      </v-card>
      <br>
      <br>
      <!-- <v-card>
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
      </v-card> -->
      <br>
      <br>
      <v-card>
        <v-card-title>一覧</v-card-title>
        <v-card-text>
          <v-data-table
            :headers="votesHeaders"
            :items="votesDesserts"
            :items-per-page="5"
            class="elevation-1"
          >
          </v-data-table>
        </v-card-text>
      </v-card>
      <!-- <v-card>
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
      </v-card> -->
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

      voteTitle: "",
      voteDescription: "",
      votedID: "",
      votesHeaders: [
        {
          text: "ID",
          align: "start",
          value: "id",
        },
        {
          text: "タイトル",
          value: "title"
        },
        {
          text: "詳細",
          value: "description",
        },
        {
          text: "作成者",
          value: "owner",
        },
        {
          text: "作成日",
          value: "created",
        },
        {
          text: "終了日",
          value: "closed"
        },
        {
          text: "賛成",
          value: "agree",
        },
        {
          text: "反対",
          value: "disagree",
        }
      ],
      votesDesserts: [],
      voteError: "",
      voteSuccess: "",
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

    // this.$axios.$post('/api/user/get_app')
    //   .then((result) => {
    //     this.appsDesserts = result
    //   })
    //   .catch((error) => {
    //     console.log(error.response.data.Msg)
    //   })

    this.$axios.$post('/api/user/vote/get')
      .then((result) => {
        console.log(result)
        this.votesDesserts = result.agendas
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
    createVote() {
      if (this.$refs.createVoteForm.validate()) {
        const params = new URLSearchParams()
        params.append('Title', this.voteTitle)
        params.append('Description', this.voteDescription)
        this.$axios.$post('/api/user/vote/create', params)
          .then((result) => {
            console.log(result)
            this.error = ""
            this.voteTitle = ""
            this.voteDescription = ""
            this.$refs.createVoteForm.resetValidation()
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
    voteAgree() {
      if (this.$refs.voteForm.validate()) {
        this.vote(true)
      }
    },
    voteDisagree() {
      if (this.$refs.voteForm.validate()) {
        this.vote(false)
      }
    },
    vote(agree) {
      const params = new URLSearchParams()
        params.append('Id', parseInt(this.votedID))
        params.append('Agree', agree)
        this.$axios.$post('/api/user/vote/vote', params)
          .then((result) => {
            console.log(result)
            this.voteError = ""
            this.voteSuccess = result.status ? "success" : "faild"
          })
          .catch((e) => {
            if (e.response) {
              this.voteError = e.response.data.Msg
            } else {
              this.voteError = "予期せぬなエラーが発生しました. 問い合わせてください."
            }
          })
    }  
  },
}
</script>