<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <v-list>
        <v-list-item @click="logout">
          <v-list-item-action>
            <v-icon>mdi-exit-run</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Logout</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
    </v-app-bar>
    <v-content>
      <v-container pa-0>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer padless app>
      &copy; {{ new Date().getFullYear() }} Kindai Univ Computer Study Group.
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      clipped: false,
      drawer: false,
      fixed: false,
      items: [
        {
          icon: 'mdi-home',
          title: 'HOME',
          to: '/user/',
        },
        {
          icon: 'mdi-record-rec',
          title: 'REC',
          to: '/user/rec',
        }
      ],
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: 'User Page'
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