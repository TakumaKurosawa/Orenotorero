<template>
  <v-app dark>
    <v-navigation-drawer v-model="drawer" app>
      <v-list>
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
    <v-app-bar app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
      <v-spacer></v-spacer>
      <div v-if="this.$store.state.auth.isAuth">
        ログインしてます
        <!--        <v-btn @click="logout">-->
        <!--          <v-icon>mdi-logout</v-icon>-->
        <!--          ログアウト-->
        <!--        </v-btn>-->
      </div>
      <v-btn v-else to="/authentication">
        <v-icon>mdi-login</v-icon>
        ログイン
      </v-btn>
    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer app>
      <span>&copy; 2019</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
      items: [
        {
          icon: 'mdi-view-headline',
          title: 'トップ',
          to: '/'
        },
        {
          icon: 'mdi-login',
          title: 'ログイン',
          to: '/authentication'
        },
        {
          icon: 'mdi-view-module',
          title: 'ボード一覧',
          to: '/home'
        }
      ],
      title: 'Orenotorero'
    }
  }
  // method() {
  //   logout(){
  //     this.$store.dispatch('auth/logout')
  //   }
  // },
}
</script>
