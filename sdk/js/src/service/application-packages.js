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

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class ApplicationPackages {
  constructor(registry) {
    this._api = registry
    autoBind(this)
  }

  async getDefaultAssociation(appId, fPort, selector) {
    const result = await this._api.GetDefaultAssociation(
      {
        routeParams: {
          'ids.application_ids.application_id': appId,
          'ids.f_port': fPort,
        },
      },
      Marshaler.selectorToFieldMask(selector),
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async setDefaultAssociation(
    appId,
    fPort,
    patch,
    mask = Marshaler.fieldMaskFromPatch(
      patch,
      this._api.SetDefaultAssociationAllowedFieldMaskPaths,
    ),
  ) {
    const result = await this._api.SetDefaultAssociation(
      {
        routeParams: {
          'default.ids.application_ids.application_id': appId,
          'default.ids.f_port': fPort,
        },
      },
      {
        default: patch,
        field_mask: Marshaler.fieldMask(mask),
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async deleteDefaultAssociation(appId, fPort) {
    const result = await this._api.DeleteDefaultAssociation({
      routeParams: {
        'application_ids.application_id': appId,
        f_port: fPort,
      },
    })

    return Marshaler.payloadSingleResponse(result)
  }

  async getAssociation(appId, fPort, deviceId, selector) {
    console.log('appId: ', appId, fPort, deviceId, selector);
    let result;
    console.log("GOT INTO GETASSOCIATION from sdk.......");
    console.log('_api: ', this._api);
    try {

      result = await this._api.GetAssociation(
        {
          routeParams: {
            'ids.end_device_ids.application_ids.application_id': appId,
            'ids.end_device_ids.device_id': deviceId,
            'ids.f_port': fPort,
          },
        },

        Marshaler.selectorToFieldMask(selector),
      )
      console.log('result: ', result);
    } catch (error) {
      console.log("Error in Marshaler.payloadSingleResponse(result).......", error);
    }

    console.log("Marshaler.payloadSingleResponse(result).......", Marshaler.payloadSingleResponse(result));

    console.log("1111");

    return Marshaler.payloadSingleResponse(result)
      .ApplicationPackages
  }

  async setAssociation(
    appId,
    deviceId,
    fPort,
    patch,
    mask = Marshaler.fieldMaskFromPatch(
      patch,
      this._api.SetAssociationAllowedFieldMaskPaths,
    ),
  ) {
    try {
      const result = await this._api.SetAssociation(
        {
          routeParams: {
            'association.ids.end_device_ids.application_ids.application_id': appId,
            'association.ids.end_device_ids.device_id': deviceId,
            'association.ids.f_port': fPort,
          },
        },
        {
          default: patch,
          field_mask: Marshaler.fieldMask(mask),
        },
      )
      console.log('result from SetAssociation: ', result);

      return Marshaler.payloadSingleResponse(result)
    } catch (error) {
      console.log("error in setAssociation: ", error);
    }
  }
}

export default ApplicationPackages
