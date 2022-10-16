import { DocsExample } from 'src/components'

import { useEffect, useState } from 'react';
import React from 'react'
import axios from 'axios'
import {
  CButton,
  CCard,
  CCardBody,
  CCardHeader,
  CCol,
  CForm,
  CFormCheck,
  CFormInput,
  CFormLabel,
  CRow,
} from '@coreui/react'

const baseURL = "http://localhost:8080/api/article";


const Layout = () => {
  const [article, setArticle] = React.useState(null);

  // React.useEffect(() => {
  //   axios.get(`${baseURL}`).then((response) => {
  //     setArticle(response.data);
  //   });
  // }, []);

  const createPost = () =>{
    axios
      .post(baseURL, {
        title: "Hello World!",
        body: "This is a new post."
      })
      .then((response) => {
        setArticle(response.data);
      });
  }

  const makeAPICall = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/article', {mode:'cors'});
      const data = await response.json();
      console.log({ data })
    }
    catch (e) {
      console.log(e)
    }
  }
  useEffect(() => {
    makeAPICall();
  }, [])

  // if (!article) return "No post!"

  return (
    <CRow>
      
        <form action={createPost} method="post">
          <CCol xs={12}>
            <CCard className="mb-4">
              <CCardHeader>
                <strong>Form</strong> <small>Post Article</small>
              </CCardHeader>
              <CCardBody>
                
                <DocsExample href="forms/layout#gutters">
                  <CForm className="row g-3">
                    <CCol md={6}>
                      <CFormLabel htmlFor="inputTitle">Title</CFormLabel>
                      <CFormInput type="text" id="inputTitle" />
                    </CCol>
                    <CCol md={6}>
                      <CFormLabel htmlFor="inputContent">Content</CFormLabel>
                      <CFormInput type="text" id="inputContent" />
                    </CCol>

                    <CCol md={6}>
                      <CFormLabel htmlFor="inputCategory">Category</CFormLabel>
                      <CFormInput type="text" id="inputCategory" />
                    </CCol>
                    <fieldset className="row mb-3">
                      <legend className="col-form-label col-sm-2 pt-0">Radios</legend>
                      <CCol sm={10}>
                        <CFormCheck
                          type="radio"
                          name="statusRadio"
                          id="radio1"
                          value="Publish"
                          label="Publish"
                          defaultChecked
                        />
                        <CFormCheck
                          type="radio"
                          name="statusRadio"
                          id="radio2"
                          value="Draft"
                          label="Draft"
                        />
                        <CFormCheck
                          type="radio"
                          name="statusRadio"
                          id="radio3"
                          value="Trash"
                          label="Trash"
                        />
                      </CCol>
                    </fieldset>
                    
                    <CCol xs={12}>
                      <CButton type="submit" >Sign in</CButton>
                    </CCol>
                  </CForm>
                </DocsExample>
              </CCardBody>
            </CCard>
          </CCol>
        </form>
      
      
      
      
    </CRow>
  )
}

export default Layout
