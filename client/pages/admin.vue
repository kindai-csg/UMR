<template>
  <v-layout column>
    <v-card>
      <v-card-title>アカウント管理画面</v-card-title>
      <v-card-text>
        <button @click="logout">Logout</button>
      </v-card-text>
    </v-card>
    <br>
    <br>
    <v-card>
      <v-card-title>登録フォーム管理</v-card-title>
      <v-card-text>
        <v-alert type="error" v-if="formError">
          {{ formError }}
        </v-alert>
        <div v-if="formUrl&&formTime">
          フォームURL: {{ formUrl }}
          <br>
          残り時間: {{ formTime }} 秒 
        </div>
        <br>
        <v-form ref="createForm">
          <v-text-field
          v-model="newFormTime"
          label="有効期限(秒)"
          :rules="[required, numberCheck]"
          />
          <v-btn text v-on:click="createForm">作成</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
    <br>
    <br>
    <v-card>
      <v-card-title>アカウント管理</v-card-title>
      <v-card-text>
        <v-alert type="error" v-if="accountError">
          {{ accountError }}
        </v-alert>
        <v-data-table
          :headers="accountsHeaders"
          :items="accountsDesserts"
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
    <br>
    <br>
    <v-card>
      <v-card-title>アカウントアクティベイト</v-card-title>
      <v-card-text>
        <v-checkbox v-model="autoActivate" label="自動アクティベイト"/>
        <v-alert type="error" v-if="activateError">
          {{ activateError }}
        </v-alert>
        <v-data-table
          :headers="activateHeaders"
          :items="activateDesserts"
          :items-per-page="5"
          class="elevation-1"
        >
          <template v-slot:item.actions="{ item }">
            <v-icon
              small
              class="mr-2"
              @click="activateItem(item)"
            >
              mdi-account-check
            </v-icon>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>
  </v-layout>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      newFormTime: 600,
      formUrl: "",
      formTime: 0,
      formError: "",
      accountError: "",
      activateError: "",
      required: value => !!value || "必須項目です",
      numberCheck: value => !isNaN(value) || "半角英数字のみで入力してください",
      autoActivate: false,

      accountsHeaders: [
        {
          text: "ユーザーID",
          align: "start",
          value: "ID",
        },
        {
          text: "学籍番号",
          value: "StudentNumber",
        },
        {
          text: "名前",
          value: "Name",
        },
        {
          text: "メールアドレス",
          value: "EmailAddress",
        },
        {
          text: "アクション",
          value: 'actions',
          sortable: false,
        }
      ],
      accountsDesserts: [],
      activateHeaders: [
        {
          text: "ユーザーID",
          align: "start",
          value: "ID",
        },
        {
          text: "アクション",
          value: 'actions',
          sortable: false,
        }
      ],
      activateDesserts: [],
    }
  },
  head() {
    return {
      title: 'アカウント管理画面',
    }
  },
  created() {
    this.$axios.$post('/api/get_token_authority')
      .then((result) => {
        if (!result.Admin) {
          this.$router.push('/user')
        }
      })
      .catch((error) => {
        this.$router.push('/')
      })
  },
  mounted() {
    // 登録フォームの取得
    this.$axios.$post('/api/admin/get_register_form')
      .then((result) => {
        this.formUrl = location.origin + "/register?token=" + result.Token
        this.formTime = result.Time
      })

    // 登録フォームの生存秒数の表示の更新 (1秒ごと)
    setInterval(() => {
      if (this.formTime > 0) {
        this.formTime -= 1
      }
    }, 1000)

    // 全アカウントの取得
    this.$axios.$post('/api/admin/get_all_accounts')
      .then((result) => {
        this.accountsDesserts = result
      })
      .catch((error) => {
        this.accountError =  error.response.data.Msg
      })

    // 全アカウントの再取得 (5秒ごと)
    setInterval(() => {
      this.$axios.$post('/api/admin/get_all_accounts')
        .then((result) => {
          this.accountsDesserts = result
        })
    }, 5000)

    // アクティベイト前のアカウントの取得
    this.$axios.$post('/api/admin/get_all_non_active_account_id')
      .then((result) => {
        this.activateDesserts = result
      })
      .catch((error) => {
        this.activateError = error.response.data.Msg
      })
    
    // アクティベイト前のアカウントの再取得 (5秒ごと)
    // もし自動アクティベイトがオンならついでにアクティベイトに投げる
    setInterval(() => {
      this.$axios.$post('/api/admin/get_all_non_active_account_id')
        .then((result) => {
          this.activateDesserts = result
          if (!this.autoActivate) {
            return
          }
          this.activateDesserts.forEach(account => {
            const params = new URLSearchParams()
            params.append('ID', account.ID)
            this.$axios.$post('/api/admin/activation', params)
              .then((result) => {
              })
          })
        })
    }, 5000)

  },
  methods: {
    /**
     * 管理画面からログアウトする(token削除)
     */
    logout() {
      this.$auth.logout()
    },
    /**
     * 登録フォームを生成する
     */
    createForm() {
      if (this.$refs.createForm.validate()) {
        const params = new URLSearchParams()
        params.append('Time', this.newFormTime)
        this.$axios.$post('/api/admin/create_register_form', params)
          .then((result) => {
            this.formUrl = location.origin + "/register?token=" + result.Token
            this.formTime = result.Time
          })
          .catch((e) => {
            if (e.response) {
              this.formError = e.response.data.Msg
            } else {
              this.formError = "予期せぬなエラーが発生しました." 
            }
          })
      }
    },
    /**
     * 指定アカウントをアクティベイトする
     */
    activateItem(item) {
      const params = new URLSearchParams()
      params.append('ID', item.ID)
      this.$axios.$post('/api/admin/activation', params)
        .then(() => {
          this.activateError = ''
          const index = this.activateDesserts.indexOf(item)
          this.activateDesserts.splice(index, 1)
        })
        .catch((error) => {
          this.activateError =  error.response.data.Msg
        })
    },
    /**
     * 指定アカウントを削除する
     */
    deleteAccount(account) {
      const params = new URLSearchParams()
      params.append('ID', account.ID)
      this.$axios.$post('/api/admin/delete_account', params)
        .then(() => {
          this.accountError = ''
          const index = this.accountsDesserts.indexOf(account)
          this.accountsDesserts.splice(index, 1)
        })
        .catch((error) => {
          this.accountError =  error.response.data.Msg
        })
    }
  }
}
</script>