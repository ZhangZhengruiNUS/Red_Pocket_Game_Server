<template>

  <v-btn
      @click="router.push({ name: 'Catalog' })"
      color="primary"
      variant="elevated">
    Back to catalog
  </v-btn>

  <div class="product">
    <div class="product-image">
      <img :src="selectedProduct.thumbnail" alt="">
    </div>
    <div class="product-details">
      <p>Brand: {{ selectedProduct.brand }}</p>
      <p>Description: {{ selectedProduct.description }}</p>
      <h2>Price: ${{ selectedProduct.price }}</h2>
      <v-btn
          variant="elevated"
          color="indigo-lighten-3"
          @click="addToCart"
      >Add to cart</v-btn>
    </div>
  </div>

 <el-table :data="tableData">
      <el-table-column prop="id" label="Id"></el-table-column>
      <el-table-column prop="title" label="Title"></el-table-column>
      <el-table-column prop="price" label="Price"></el-table-column>
      <el-table-column>
        <template #default="{row}">
          <el-button @click="editRow(row)">Edit</el-button>
        </template>
      </el-table-column>
    </el-table>
  <el-dialog :visible.sync="dialogVisible">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="id" prop="id">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="Age" prop="title">
          <el-input v-model.number="form.title"></el-input>
        </el-form-item>
        <el-form-item label="Address" prop="price">
          <el-input v-model="form.price"></el-input>
        </el-form-item>
      </el-form>
 
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="submitForm">Save</el-button>
      </div>
    </el-dialog>

<div class="update">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>Update Product</span>
            </div>
            <el-tabs
                v-model="currentIndex"
                stretch
                @tab-click="handleTabsClick"
            >

         <el-tab-pane label="登陆" name="update">
                    <el-form
                        :model="updateForm"
                        :rules="rules"
                        status-icon
                        ref="updateForm"
                    >
                        <el-form-item
                            label="Title"
                            label-width="80px"
                            prop="title"
                        >
                            <el-input
                                type="text"
                                v-model="updateForm.title"
                                class="border"
                            />
                        </el-form-item>
                        <el-form-item
                            label="Price"
                            label-width="80px"
                            prop="price"
                        >
                            <el-input
                                type="text"
                                v-model="updateForm.price"
                                class="border"
                            />
                        </el-form-item>
                        <el-form-item>
                            <el-button
                                type="primary"
                                @click="submitForm('updateForm')"
                                >提交</el-button
                            >
                        </el-form-item>
                    </el-form>
                </el-tab-pane>
             </el-tabs>
       </el-card>
    </div>
 </template> 
<script>
  import { defineComponent } from "vue";
  import { ref } from 'vue';
  import { ElTable, ElTableColumn, ElButton, ElDialog, ElForm, ElFormItem, ElInput } from 'element-plus';
  import { computed } from "vue";
  import { productsStore } from "@/stores/products";
  import { useRoute, useRouter } from "vue-router";
  import { mapMutations } from "vuex";
  export default {
  components: {
    ElTable,
    ElTableColumn,
    ElButton,
    ElDialog,
    ElForm,
    ElFormItem,
    ElInput,
    name: 'ProductDetails'
 
  },

data() {

 const form = ref({});
  const dialogVisible = ref(true);
  const store = productsStore()
  const router = useRouter()
  const route = useRoute()
  const tableData = ref(store.products)
  console.log(store.products)

        // 验证规则
        var validateUserName = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("Please enter title"));
            } 
            else {
                callback();
            }
        };
        var validatePassWord = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("Please enter price"));
            } else {
                callback();
            }
        };
        var validateConfigurePassword = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请输入密码"));
            } else if (value !== this.registerForm.password) {
                callback(new Error("两次输入密码不一致!"));
            } else {
                callback();
            }
        };
         
  const addToCart = () => {
    store.addToCart(selectedProduct.value)
    router.push({ name: 'CartView' })
  }
const selectedProduct = computed(() => {
    return store.products.find((item) => item.id === Number(route.params.id))
  })
        return {
            currentIndex: "update",
            updateForm: {
                title: "",
                price: "",
            },
            registerForm: {
                title: "",
                price: "",
                configurePassword: "",
                email: "",
            },
            activeTab: "update",
            selectedProduct,
            rules: {
                title: [
                    {
                        validator: validateUserName,
                        trigger: "blur",
                    },
                ],
                price: [
                    {
                        validator: validatePassWord,
                        trigger: "blur",
                    },
                ],
                configurePassword: [
                    {
                        validator: validateConfigurePassword,
                        trigger: "blur",
                    },
                ],
            },
        };
    },
    methods: {
        ...mapMutations("update", ["setUser"]),
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    if (this.activeTab === "update") {
                        // 登陆
this.selectedProduct.value.title=this.updateForm.title
this.selectedProduct.value.price=this.updateForm.price
                    }

                } else {
                    return;
                }
            });
        },
        handleTabsClick(tab) {
            this.activeTab = tab.name;
        },
    },
};

  
  
</script>

<style scoped lang="less">
.product {
  display: flex;
  margin-top: 50px;
}

.product-image {
  margin-right: 24px;
}
.update {
    width: 300px;
    margin: 0 auto;
    .box-card {
        width: 10px;
        margin: 10px auto;
        
    }
}
.border{
    border: 1px solid;
}
</style>