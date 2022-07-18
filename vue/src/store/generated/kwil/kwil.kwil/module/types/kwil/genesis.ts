/* eslint-disable */
import { Params } from "../kwil/params";
import { Databases } from "../kwil/databases";
import { Ddl } from "../kwil/ddl";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "kwil.kwil";

/** GenesisState defines the kwil module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  databasesList: Databases[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  ddlList: Ddl[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.databasesList) {
      Databases.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.ddlList) {
      Ddl.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.databasesList = [];
    message.ddlList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.databasesList.push(Databases.decode(reader, reader.uint32()));
          break;
        case 3:
          message.ddlList.push(Ddl.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.databasesList = [];
    message.ddlList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.databasesList !== undefined && object.databasesList !== null) {
      for (const e of object.databasesList) {
        message.databasesList.push(Databases.fromJSON(e));
      }
    }
    if (object.ddlList !== undefined && object.ddlList !== null) {
      for (const e of object.ddlList) {
        message.ddlList.push(Ddl.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.databasesList) {
      obj.databasesList = message.databasesList.map((e) =>
        e ? Databases.toJSON(e) : undefined
      );
    } else {
      obj.databasesList = [];
    }
    if (message.ddlList) {
      obj.ddlList = message.ddlList.map((e) => (e ? Ddl.toJSON(e) : undefined));
    } else {
      obj.ddlList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.databasesList = [];
    message.ddlList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.databasesList !== undefined && object.databasesList !== null) {
      for (const e of object.databasesList) {
        message.databasesList.push(Databases.fromPartial(e));
      }
    }
    if (object.ddlList !== undefined && object.ddlList !== null) {
      for (const e of object.ddlList) {
        message.ddlList.push(Ddl.fromPartial(e));
      }
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
