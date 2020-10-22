/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
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
    EntRoomInfoEdges,
    EntRoomInfoEdgesFromJSON,
    EntRoomInfoEdgesFromJSONTyped,
    EntRoomInfoEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoomInfo
 */
export interface EntRoomInfo {
    /**
     * INFOBED holds the value of the "INFOBED" field.
     * @type {number}
     * @memberof EntRoomInfo
     */
    iNFOBED?: number;
    /**
     * INFOOFFICEDE holds the value of the "INFOOFFICEDE" field.
     * @type {number}
     * @memberof EntRoomInfo
     */
    iNFOOFFICEDE?: number;
    /**
     * INFOREFRIGERAT holds the value of the "INFOREFRIGERAT" field.
     * @type {number}
     * @memberof EntRoomInfo
     */
    iNFOREFRIGERAT?: number;
    /**
     * INFOWARDROB holds the value of the "INFOWARDROB" field.
     * @type {number}
     * @memberof EntRoomInfo
     */
    iNFOWARDROB?: number;
    /**
     * 
     * @type {EntRoomInfoEdges}
     * @memberof EntRoomInfo
     */
    edges?: EntRoomInfoEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntRoomInfo
     */
    id?: number;
}

export function EntRoomInfoFromJSON(json: any): EntRoomInfo {
    return EntRoomInfoFromJSONTyped(json, false);
}

export function EntRoomInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'iNFOBED': !exists(json, 'INFOBED') ? undefined : json['INFOBED'],
        'iNFOOFFICEDE': !exists(json, 'INFOOFFICEDE') ? undefined : json['INFOOFFICEDE'],
        'iNFOREFRIGERAT': !exists(json, 'INFOREFRIGERAT') ? undefined : json['INFOREFRIGERAT'],
        'iNFOWARDROB': !exists(json, 'INFOWARDROB') ? undefined : json['INFOWARDROB'],
        'edges': !exists(json, 'edges') ? undefined : EntRoomInfoEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntRoomInfoToJSON(value?: EntRoomInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'INFOBED': value.iNFOBED,
        'INFOOFFICEDE': value.iNFOOFFICEDE,
        'INFOREFRIGERAT': value.iNFOREFRIGERAT,
        'INFOWARDROB': value.iNFOWARDROB,
        'edges': EntRoomInfoEdgesToJSON(value.edges),
        'id': value.id,
    };
}


