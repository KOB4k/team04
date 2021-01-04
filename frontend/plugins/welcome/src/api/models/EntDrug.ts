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
    EntDrugEdges,
    EntDrugEdgesFromJSON,
    EntDrugEdgesFromJSONTyped,
    EntDrugEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrug
 */
export interface EntDrug {
    /**
     * DrugName holds the value of the "DrugName" field.
     * @type {string}
     * @memberof EntDrug
     */
    drugName?: string;
    /**
     * Howto holds the value of the "Howto" field.
     * @type {string}
     * @memberof EntDrug
     */
    howto?: string;
    /**
     * Property holds the value of the "Property" field.
     * @type {string}
     * @memberof EntDrug
     */
    property?: string;
    /**
     * 
     * @type {EntDrugEdges}
     * @memberof EntDrug
     */
    edges?: EntDrugEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDrug
     */
    id?: number;
}

export function EntDrugFromJSON(json: any): EntDrug {
    return EntDrugFromJSONTyped(json, false);
}

export function EntDrugFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrug {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'drugName': !exists(json, 'DrugName') ? undefined : json['DrugName'],
        'howto': !exists(json, 'Howto') ? undefined : json['Howto'],
        'property': !exists(json, 'Property') ? undefined : json['Property'],
        'edges': !exists(json, 'edges') ? undefined : EntDrugEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDrugToJSON(value?: EntDrug | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'DrugName': value.drugName,
        'Howto': value.howto,
        'Property': value.property,
        'edges': EntDrugEdgesToJSON(value.edges),
        'id': value.id,
    };
}


