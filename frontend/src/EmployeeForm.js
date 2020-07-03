import react, * as React from "react";
import { Page, Grid } from "tabler-react";
import SiteWrapper from "./SiteWrapper.react";
import { Button, Form, FormGroup, Label, Input } from 'reactstrap';
import { withFormik } from 'formik';

const EmployeeForm = ({ values, handleChange, handleSubmit, errors, touched, isSubmitting }) => {
  return (
    <SiteWrapper>
      <Page.Card
            title="Employee Registration"
        ></Page.Card>
        <Grid.Col md={6} lg={6} className="align-self-center">
        <Form onSubmit={handleSubmit}>
          <FormGroup>
            {touched.username && errors.username && <p className="red">{errors.username}</p>}
            <Label for="username">Email</Label>
            <Input 
              type="text" 
              name="username"
              value={values.username}
              onChange={handleChange}
              id="username" 
              placeholder="Your user" 
            />
          </FormGroup>
          <FormGroup>
            {touched.password && errors.password && <p className="red">{errors.password}</p>}
            <Label for="password">Password</Label>
            <Input 
              type="password" 
              name="password"
              value={values.password}
              onChange={handleChange}
              id="password" 
              placeholder="Your secret password" 
            />
          </FormGroup>
          <Button color="primary" disabled={isSubmitting}>Submit</Button>
        </Form>
    </Grid.Col>
    </SiteWrapper>
  );
}

const FormikApp = withFormik({
  mapPropsToValues({ username, password }) {
    return { username, password }
  },
  // validationSchema: Yup.object().shape({
  //   username: Yup.string().email().required('Email is required'),
  //   password: Yup.string().min(4, 'Password must be 4 characters or longer').required()
  // }),
  handleSubmit(values, { props, resetForm, setErrors, setSubmitting }) {
    alert(JSON.stringify(values))
    // if (values.username === 'adeel@io.com') {
    //   setErrors({ username: 'This is a dummy procedure error' });
    // } else {
    //   props.onSubmit(values);
    //   resetForm();
      // setSubmitting(true);
    // }
  }
})(EmployeeForm);

export default FormikApp
// const FormElements = () => (
//   <SiteWrapper>
//          <Page.Card
//             title="Employee Registration"
//             // RootComponent={Form}
//         ></Page.Card>
//         <Grid.Col md={6} lg={6} className="align-self-center">
//   <div>
//     <Formik
//     initialValues={{ email: "" }}
//     onSubmit={async values => {
//       await new Promise(resolve => setTimeout(resolve, 500));
//       alert(JSON.stringify(values, null, 2));
//     }}
//     >
//     {props => {
//         const {
//           values,
//           touched,
//           errors,
//           dirty,
//           isSubmitting,
//           handleChange,
//           handleBlur,
//           handleSubmit,
//           handleReset
//         } = props;
//         return (
//           <form onSubmit={handleSubmit}>
//             <label htmlFor="id" style={{ display: "block" }}>
//               ID
//             </label>
//             <input
//               id="id"
//               placeholder="Employee ID"
//               type="text"
//               onChange={handleChange}
//               onBlur={handleBlur}
//             />

//             <label htmlFor="name" style={{ display: "block" }}>
//               Name
//             </label>
//             <input
//               id="name"
//               placeholder="Employee Name"
//               type="text"
//               onChange={handleChange}
//               onBlur={handleBlur}
//             />

//             <label htmlFor="job_role" style={{ display: "block" }}>
//               Job Role
//             </label>
//             <select
//               name="job_role"
//               value={values.job_role}
//               onChange={handleChange}
//               onBlur={handleBlur}
//               style={{ display: 'block' }}
//             >
//               <option value="red" label="red" />
//               <option value="blue" label="blue" />
//               <option value="green" label="green" />
//             </select>

//             <button
//               type="button"
//               className="outline"
//               onClick={handleReset}
//               disabled={!dirty || isSubmitting}
//             >
//               Reset
//             </button>
//             <button type="submit" disabled={isSubmitting}>
//               Submit
//             </button>
//           </form>
//         );
//       }}
//   </Formik>
//   </div>
//   </Grid.Col>
//   </SiteWrapper>
// );

// export default FormElements
