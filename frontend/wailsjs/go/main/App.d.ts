// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {utils} from '../models';
import {mtp} from '../models';

export function AddGame(arg1:Array<utils.Game>):Promise<utils.Msg>;

export function AddTask(arg1:number,arg2:string):Promise<utils.Msg>;

export function CancelTask(arg1:number):Promise<utils.Msg>;

export function DeleteGame(arg1:Array<number>):Promise<utils.Msg>;

export function DeleteWorkOrder(arg1:Array<number>):Promise<utils.Msg>;

export function GetAnnouncement():Promise<utils.Msg>;

export function GetDeviceList():Promise<Array<mtp.Device>>;

export function GetGameList(arg1:string):Promise<Array<utils.Game>>;

export function GetOrderGameList(arg1:number):Promise<Array<utils.WorkOrderGameList>>;

export function GetWorkOrder(arg1:utils.WorkOrderQuery):Promise<any>;

export function UpdateAnnouncement(arg1:Array<string>):Promise<utils.Msg>;

export function UpdateGame(arg1:utils.Game):Promise<utils.Msg>;