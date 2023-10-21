// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { createApplicationApiKey } from '@console/store/actions/api-keys'
import DataSheet from '@ttn-lw/components/data-sheet'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import React, { useCallback, useEffect, useRef, useState } from 'react'
import { defineMessages } from 'react-intl'
import { useDispatch } from 'react-redux'
import { Link, useParams } from 'react-router-dom'
import ModalButton from '@ttn-lw/components/button/modal-button'
import Form from '@ttn-lw/components/form'
import Input from '@ttn-lw/components/input'
import SubmitBar from '@ttn-lw/components/submit-bar'
import SubmitButton from '@ttn-lw/components/submit-button'
import Yup from '@ttn-lw/lib/yup'
import Button from '@ttn-lw/components/button'

const m = defineMessages({
  tokenDescription: 'Fuota access token as configured within FUOTA.IO',
  deleteWarning: 'Are you sure you want to delete the FUOTA.IO token?',
  connectionCredentials: 'Fuota API Key',
})

const FuotaIoForm = () => {
  const [error, setError] = useState('')

  const [apiKey, setApiKey] = useState('')

  const [loading, setLoading] = useState(false)

  const formRef = useRef(null)

  const validationSchema = Yup.object()
    .shape({
      data: Yup.object().shape({
        apiKey: Yup.string().default('').required(sharedMessages.validateRequired),
      }),
    })
    .noUnknown()

  const handleSubmit = values => {
    setLoading(true)
    setTimeout(() => {
      setApiKey(values.data.apiKey)
      setLoading(false)
    }, 1000)
  }

  const handleDelete = () => {
    setApiKey('')
    formRef.current.setFieldValue('data.apiKey', '')
  }

  const initialValues = validationSchema.cast({})

  return (
    <Form
      error={error}
      validationSchema={validationSchema}
      initialValues={initialValues}
      onSubmit={handleSubmit}
      formikRef={formRef}
    >
      <Form.Field
        component={Input}
        title={sharedMessages.apiKey}
        description={m.tokenDescription}
        name="data.apiKey"
        sensitive
        required
      />
      <SubmitBar>
        <Button primary type="submit" disabled={apiKey} busy={loading}>
          Set Token
        </Button>
        {apiKey && (
          <ModalButton
            type="button"
            icon="delete"
            message={sharedMessages.tokenDelete}
            modalData={{
              message: m.deleteWarning,
            }}
            onApprove={handleDelete}
            danger
            naked
          />
        )}
      </SubmitBar>
    </Form>
  )
}

export default FuotaIoForm
