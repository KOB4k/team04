/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCategoryEdges
 */
export interface EntCategoryEdges {
    /**
     * Patient holds the value of the patient edge.
     * @type {Array<EntPatient>}
     * @memberof EntCategoryEdges
     */
    patient?: Array<EntPatient>;
}

export function EntCategoryEdgesFromJSON(json: any): EntCategoryEdges {
    return EntCategoryEdgesFromJSONTyped(json, false);
}

export function EntCategoryEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCategoryEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patient': !exists(json, 'patient') ? undefined : ((json['patient'] as Array<any>).map(EntPatientFromJSON)),
    };
}

export function EntCategoryEdgesToJSON(value?: EntCategoryEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'patient': value.patient === undefined ? undefined : ((value.patient as Array<any>).map(EntPatientToJSON)),
    };
}


