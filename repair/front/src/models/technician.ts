import request from '~/global/request'

export default class Technician {
    static grades = [' ', '건축초급기술자', '건축중급기술자', '건축고급기술자', '건축특급기술자']    
    static gradeTypes = ['', 'info', 'success', 'warning', 'danger']

    static getGradeType(value: number) {
        return this.gradeTypes[value]
    }

    static getGrade(value: number) {
        return this.grades[value]
    }

    static async insert(item: any) {
        const res = await request({
            method: 'POST',
            url: '/api/technician',
            data: item
        })

        return res
    }

    static async update(item: any) {
        const res = await request({
            method: 'PUT',
            url: '/api/technician',
            data: item
        })

        return res
    }

    static async remove(item: any) {
        const res = await request({
            method: 'DELETE',
            url: '/api/technician',
            data: item
        })

        return res
    }

    static async find(params: any) {
        const res = await request({
            method: 'GET',
            url: '/api/technician',
            params: params
        })

        return res
    }

    static async get(id: number) {
        const res = await request({
            method: 'GET',
            url: `/api/technician/${id}`
        })

        return res
    }
}
