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
    EntBook,
    EntBookFromJSON,
    EntBookFromJSONTyped,
    EntBookToJSON,
    EntUserStatus,
    EntUserStatusFromJSON,
    EntUserStatusFromJSONTyped,
    EntUserStatusToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * USERBOOK holds the value of the USER_BOOK edge.
     * @type {Array<EntBook>}
     * @memberof EntUserEdges
     */
    userbook?: Array<EntBook>;
    /**
     * 
     * @type {EntUserStatus}
     * @memberof EntUserEdges
     */
    useruserstatus?: EntUserStatus;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'userbook': !exists(json, 'USERBOOK') ? undefined : ((json['USERBOOK'] as Array<any>).map(EntBookFromJSON)),
        'useruserstatus': !exists(json, 'USERUSERSTATUS') ? undefined : EntUserStatusFromJSON(json['USERUSERSTATUS']),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'userbook': value.userbook === undefined ? undefined : ((value.userbook as Array<any>).map(EntBookToJSON)),
        'useruserstatus': EntUserStatusToJSON(value.useruserstatus),
    };
}


