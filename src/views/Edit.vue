<template>
  <div>
    <el-table :data="tableData">
      <el-table-column prop="name" label="Name"></el-table-column>
      <el-table-column prop="age" label="Age"></el-table-column>
      <el-table-column prop="address" label="Address"></el-table-column>
      <el-table-column>
        <template #default="{row}">
          <el-button @click="editRow(row)">Edit</el-button>
        </template>
      </el-table-column>
    </el-table>
 
    <el-dialog :visible.sync="dialogVisible">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="Name" prop="name">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="Age" prop="age">
          <el-input v-model.number="form.age"></el-input>
        </el-form-item>
        <el-form-item label="Address" prop="address">
          <el-input v-model="form.address"></el-input>
        </el-form-item>
      </el-form>
 
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="submitForm">Save</el-button>
      </div>
    </el-dialog>
  </div>
</template>
 
<script>
import { ref } from 'vue';
import { ElTable, ElTableColumn, ElButton, ElDialog, ElForm, ElFormItem, ElInput } from 'element-plus';
 
export default {
  components: {
    ElTable,
    ElTableColumn,
    ElButton,
    ElDialog,
    ElForm,
    ElFormItem,
    ElInput,
  },
  setup() {
    const tableData = ref([
      {
        name: 'John',
        age: 30,
        address: 'New York',
      },
      {
        name: 'Jane',
        age: 25,
        address: 'San Francisco',
      },
      {
        name: 'Bob',
        age: 40,
        address: 'Dallas',
      },
    ]);
 
    const form = ref({});
    const dialogVisible = ref(false);
 
    const rules = {
      name: [
        { required: true, message: 'Please input name', trigger: 'blur' },
      ],
      age: [
        { required: true, message: 'Please input age', trigger: 'blur' },
        { type: 'number', message: 'Age must be a number', trigger: 'blur' },
      ],
      address: [
        { required: true, message: 'Please input address', trigger: 'blur' },
      ],
    };
 
    const editRow = (row) => {
      form.value = { ...row };
      dialogVisible.value = true;
    };
 
    const submitForm = () => {
      const formRef = this.$refs.form;
      formRef.validate((valid) => {
        if (valid) {
          const dataIndex = tableData.value.indexOf(form.value);
          const tableDataCopy = [...tableData.value];
          tableDataCopy.splice(dataIndex, 1, form.value);
          tableData.value = tableDataCopy;
          dialogVisible.value = false;
        }
      });
    };
 
    return {
      tableData,
      form,
      dialogVisible,
      rules,
      editRow,
      submitForm,
    };
  },
};
</script>
