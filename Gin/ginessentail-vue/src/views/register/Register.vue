<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="user.name"
                type="text"
                placeholder="请输入您的名称(选填)"
                required
              ></b-form-input>
            </b-form-group>
            <br />
            <b-form-group label="手机号">
              <b-form-input
                v-model="$v.user.telephone.$model"
                type="number"
                placeholder="请输入您的手机号"
                :state="validateState('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')"
                >手机号不合法</b-form-invalid-feedback
              >
            </b-form-group>
            <br />
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="请输入您的密码"
                :state="validateState('password')"
              ></b-form-input>
            </b-form-group>
            <b-form-invalid-feedback :state="validateState('password')"
              >密码必须不小于6位</b-form-invalid-feedback
            >
            <br />
            <b-form-group>
              <b-button
                variant="outline-primary"
                style="width: 100%"
                @click="register"
                >注册</b-button
              >
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength, maxLength } from "vuelidate/lib/validators";
import customValidator from "@/helper/validator";
import { mapActions } from "vuex";
import userService from "@/service/userService";
import storageService from "@/service/storageService";
export default {
  data() {
    return {
      user: {
        name: "",
        telephone: "",
        password: "",
      },
    };
  },
  validations: {
    user: {
      name: {
        maxLength: maxLength(30),
      },
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    // ...mapMutations("userModule", ["SET_TOKEN", "SET_USER_INFO"]),
    ...mapActions("userModule", { userRegister: "register" }),
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      userService
        .register(this.user)
        .then((res) => {
          //保存token
          console.log(res.data.data.token);
          storageService.set(storageService.USER_TOKEN, res.data.data.token);
          // this.$router.replace({ name: "register" });

          userService.info().then((res) => {
            //保存用户信息
            // console.log(res.data.data.user);
            storageService.set(
              storageService.USER_INFO,
              JSON.stringify(res.data.data.user)
            );
            //跳转主页
            this.$router.replace({ name: "Home" });
          });
        })
        .catch((err) => {
          console.log("err:", err.response.data.msg);
          this.$bvToast.toast("数据验证错误", {
            title: err.response.data.msg,
            variant: "danger",
            solid: true,
          });
        });
    },
  },
};
</script>

<style scoped>
</style>