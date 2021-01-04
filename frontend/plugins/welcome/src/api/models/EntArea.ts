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
    EntAreaEdges,
    EntAreaEdgesFromJSON,
    EntAreaEdgesFromJSONTyped,
    EntAreaEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntArea
 */
export interface EntArea {
    /**
     * AreaName holds the value of the "AreaName" field.
     * @type {string}
     * @memberof EntArea
     */
    areaName?: string;
    /**
     * 
     * @type {EntAreaEdges}
     * @memberof EntArea
     */
    edges?: EntAreaEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntArea
     */
    id?: number;
}

export function EntAreaFromJSON(json: any): EntArea {
    return EntAreaFromJSONTyped(json, false);
}

export function EntAreaFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntArea {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'areaName': !exists(json, 'AreaName') ? undefined : json['AreaName'],
        'edges': !exists(json, 'edges') ? undefined : EntAreaEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntAreaToJSON(value?: EntArea | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'AreaName': value.areaName,
        'edges': EntAreaEdgesToJSON(value.edges),
        'id': value.id,
    };
}


