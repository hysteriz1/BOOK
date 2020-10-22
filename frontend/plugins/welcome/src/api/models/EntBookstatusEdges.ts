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
} from './';

/**
 * 
 * @export
 * @interface EntBookstatusEdges
 */
export interface EntBookstatusEdges {
    /**
     * BOOKSTATUSBOOK holds the value of the BOOKSTATUS_BOOK edge.
     * @type {Array<EntBook>}
     * @memberof EntBookstatusEdges
     */
    bookstatusbook?: Array<EntBook>;
}

export function EntBookstatusEdgesFromJSON(json: any): EntBookstatusEdges {
    return EntBookstatusEdgesFromJSONTyped(json, false);
}

export function EntBookstatusEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBookstatusEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bookstatusbook': !exists(json, 'BOOKSTATUSBOOK') ? undefined : ((json['BOOKSTATUSBOOK'] as Array<any>).map(EntBookFromJSON)),
    };
}

export function EntBookstatusEdgesToJSON(value?: EntBookstatusEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bookstatusbook': value.bookstatusbook === undefined ? undefined : ((value.bookstatusbook as Array<any>).map(EntBookToJSON)),
    };
}


