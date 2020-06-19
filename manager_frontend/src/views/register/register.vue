<template>
  <div class="login-wrapper">
    <h1 class="login-title">管理后台注册</h1>
    <el-form ref="login" :model="formItem" :rules="rules">
      <el-form-item prop="username">
        <el-input v-model="formItem.name" placeholder="用户名">
          <el-button slot="prepend" icon="el-icon-user"></el-button>
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          type="password"
          v-model="formItem.password"
          placeholder="请输入密码"
        >
          <el-button slot="prepend" icon="el-icon-lock"></el-button>
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="register">注册</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { mapActions } from "vuex";

export default {
  name: "register",
  data() {
    return {
      formItem: {
        name: "",
        password: ""
      },
      rules: {
        name: [
          {
            required: true,
            message: "请输入用户名",
            trigger: "blur"
          }
        ],
        password: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur"
          },
          {
            type: "string",
            min: 6,
            message: "长度至少为6个字符",
            trigger: "blur"
          }
        ]
      }
    };
  },
  methods: {
    // ...mapMutations('userModule',['SET_TOKEN','SET_USERINFO']),
    ...mapActions("userModule", { userRegister: "register" }),
    // handleSubmit(name) {
    //   this.$refs[name].validate(valid => {
    //     if (valid) {
    //       this.register();
    //     } else {
    //       this.$message.error("格式错误");
    //     }
    //   });
    // },
    register() {
      // 验证数据
      // this.$v.user.$touch(); // 即使表单为空也会报错
      // if (this.$v.user.$anyError) {
      //   // 表单验证出错直接返回
      //   return;
      // }
      // 请求
      this.userRegister(this.formItem)
        .then(() => {
          this.$router.replace({ path: "/admin" });
          this.$message("注册成功")
        })
        .catch(err => {
          this.$message.error("登录失败！" + err.response.data.msg)
        });
    }
  }
};
</script>
<style lang="scss" scoped>
.login-wrapper {
  box-sizing: border-box;
  position: absolute;
  top: 100px;
  left: 50%;
  padding: 40px 40px;
  width: 500px;
  height: 350px;
  border-radius: 4px;
  margin-left: -250px;
  box-shadow: 0 0 10px #ccc;
  .login-title {
    margin-bottom: 40px;
  }
  .login-input {
    margin: 0 auto;
  }
}
</style>
