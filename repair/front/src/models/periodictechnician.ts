import request from '~/global/request'

export default class Periodictechnician {
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
            url: '/api/periodictechnician',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/periodictechnician',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodictechnician',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/periodictechnician',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/periodictechnician/${id}`
        })

        return res
    }

    static async deleteByPeriodic(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/periodictechnician/deletebyperiodic',
            data: item
        })

        return res
    }
}
