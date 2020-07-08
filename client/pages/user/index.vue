<template>
  <v-layout column>
    <h1>test</h1>
  </v-layout>
</template>

<script>
export default {
  layout: 'user',
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
  }
}
</script>