import request from '~/global/request'

export default class Detailtechnician {
    static types = [' ', '책임기술자', '참여기술자']    
    static typeTypes = ['', 'danger', 'warning']

    static getTypeType(value: number) {
        return this.typeTypes[value]
    }

    static getType(value: number) {
        return this.types[value]
    }

    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/detailtechnician',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/detailtechnician',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/detailtechnician',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/detailtechnician',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/detailtechnician/${id}`
        })

        return res
    }

    static async deleteByDetail(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/detailtechnician/deletebydetail',
            data: item
        })

        return res
    }
}
