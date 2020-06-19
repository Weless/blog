<template>
  <div class="admin-base">
    <el-menu :default-active="activeIndex" mode="horizontal" router>
      <el-menu-item index="/admin/management">
        <i class="el-icon-document"></i>
        <span>内容管理</span>
      </el-menu-item>
      <!-- <el-menu-item index="/admin/imageManagement">
        <i class="el-icon-picture-outline"></i>
        <span>图片管理</span>
      </el-menu-item> -->
      <el-menu-item index="/admin/categoryMangement">
        <i class="el-icon-milk-tea"></i>
        <span>分类管理</span>
      </el-menu-item>
      <el-menu-item index="/admin/newArticle">
        <i class="el-icon-plus"></i>
        <span>新增文章</span>
      </el-menu-item>
      <div class="user">
        <span class="user-name">Hi,{{ userInfo.name }}</span>
        <span class="log-out" @click="goHome">退出</span>
      </div>
    </el-menu>
    <router-view></router-view>
  </div>
</template>

<script>
import {mapState, mapActions} from 'vuex'
export default {
  computed: mapState({
      userInfo:(state) => state.userModule.userInfo,
      activeIndex() {
        return this.$route.path;
      }
  }),
  methods: {
    ...mapActions('userModule',['logout']),
    goHome(){
      this.logout()
      this.$router.replace({path:"/login"})
    }
  }
};
</script>

<style lang="scss" scoped>
.admin-base {
  .user {
    float: right;
    padding-top: 30px;
    padding-right: 40px;
    .log-out {
      margin-left: 10px;
      text-decoration: underline;
      cursor: pointer;
    }
  }
}
</style>
