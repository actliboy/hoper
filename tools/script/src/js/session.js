

function erp() {
    const sess = {
        type: 1,
        platformType: 101,
        thirdCompId: 0,
        compId: 10001,
        filterCompIds: [10001],
        employeeId: 2006,
        employeeName: 'sale2',
        piId: "1234567"
    };
    console.log(Buffer.from(JSON.stringify(sess)).toString('base64'));
}



function crm() {
    const sess = {
        userId: 200,
        userName: '',
        userRealName: '',
        clientIp: '',
    };
    console.log(Buffer.from(JSON.stringify(sess)).toString('base64'));
}

function escape() {
    const token = "eyJhY2NvdW50VHlwZSI6MCwiY29tcElkIjoxLCJjb21wTmFtZSI6Iua3seWcs+W4guWNjuWuh+iur+enkeaKgOaciemZkOWFrOWPuCIsImRlcHRJZCI6MTkzLCJkZXB0TmFtZSI6IkVSUOiZmuaLn+mDqOmXqCIsImVtcGxveWVlSWQiOjYsImVtcGxveWVlTmFtZSI6IuW5s+WPsOaWuSIsImVuZ2xpc2hOYW1lIjoicGxhdGZvcm0iLCJmaWx0ZXJDb21wSWRzIjpbMV0sImZpbHRlckRlcHRJZHMiOltdLCJmaWx0ZXJJZHMiOls2XSwicGhvbmUiOiIxNjY3NTUyMzIwMiIsInBsYXRmb3JtVHlwZSI6MTAwLCJyb2xlQ29kZUxpc3QiOlsiMyJdLCJ0aGlyZENvbXBJZCI6MCwidHlwZSI6M30="
    console.log(JSON.parse(Buffer.from(token, 'base64').toString()));
}

function encrypt() {
    const data = {"userId":2028,"userName":"yi","deptId":923,"deptName":" ","compId":17324,"compName":" ","englishName":"yi","roleCodeList":[" "],"operatorId":10001,"platformType":1,"thirdCompId":23,"systemVersion":0,"isTrial":0}
    console.log(Buffer.from(JSON.stringify(data)).toString('base64'));
}

function openErp(){
    const sess = {
        piId:"10001",
        employeeId: 2,
        employeeName: "xxx",
        deptId: 1,
        deptName: "xxx",
        compId: 5,
        compName: "xxx",
        englishName: "xxx",
        phone: "xxxxx",
        roleCodeList:["20","21","22"],
        platformType:103,
        accountType: 4,
        thirdCompId:10010,
        thirdCompCode: "",
        source:0,
        systemVersion:0,
        isTrial:0,
        type: 1,
        filterIds: [1,2,3],
        filterDeptIds: [],
        filterCompIds: []
    };
    console.log(Buffer.from(JSON.stringify(sess)).toString('base64'));
}

const all  = "e30="

console.log(Buffer.from('{}').toString('base64'));
escape()