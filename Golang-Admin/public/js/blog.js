console.log("Hello World")

let dataBlog = [];
function addBlog(event) {
  event.preventDefault();

  let title = document.getElementById("input-tittle").value;
  let content = document.getElementById("input-description").value;
  let image = document.getElementById("input-image").files;
  
  let start = new Date (document.getElementById("input-start-date").value);
  let end = new Date (document.getElementById("input-end-date").value);
  let dates = Math.abs(end - start); 
  let months = Math.floor(dates / (1000 * 60 * 60 * 24 * 30)); 
  let days = Math.floor(dates / (1000 * 60 * 60 * 24)) % 30;

  const robot = '<img class="img-icon" src="public/Image/robot.png">';
  const animal = '<img class="img-icon" src="public/Image/animal.png">';
  const demon = '<img class="img-icon" src="public/Image/demon.png">';
  const human = '<img class="img-icon" src="public/Image/human.png">';

  let checkRobot = document.getElementById("check-robot").checked ? robot: "";
  let checkAnimal = document.getElementById("check-animal").checked ? animal: "";
  let checkDemon = document.getElementById("check-demon").checked ? demon: "";
  let checkHuman = document.getElementById("check-human").checked ? human: "";

  image = URL.createObjectURL(image[0]);
  console.log(image);

  let blog = {
    title,
    months,
    days,
    content,
    image,
    checkAnimal,
    checkDemon,
    checkRobot,
    checkHuman,
    postAt: new Date(),
    author: "Alexandria",
  };

  dataBlog.push(blog);
  console.log(dataBlog);

  renderBlog();
}

function renderBlog() {
  document.getElementById("content").innerHTML = "";
  for (let index = 0; index < dataBlog.length; index++) {
    document.getElementById("content").innerHTML += `
    <div class="pt-5">
    <div id="contents" class="container responsive border">
        <div>
            <div class="blog-image margin pt-4 ps-4">
                <img src="${dataBlog[index].image}" alt="blog_img"/>
            </div>
            <div class="pt-3">
                <div class="text-box fs-6">
                    <h1 class="fs-5">
                        <a 
                        style="color: red;" 
                        href="blog-detail.html" 
                        target="_blank">${dataBlog[index].title}</a>
                    </h1>
                    <div style="color: gray; font-size: 15px;">Duration :</div>
                <div> ${dataBlog[index].months} Bulan, ${dataBlog[index].days} Hari </div>
                    <div class="detail-blog-content">
                    ${getFullTime(dataBlog[index].postAt)} | ${dataBlog[index].author}
                    </div>
                    <div>Ability</div>
                    <div>
                    ${dataBlog[index].checkAnimal}
                    ${dataBlog[index].checkDemon}
                    ${dataBlog[index].checkHuman}
                    ${dataBlog[index].checkRobot}
                    </div>
                    <p>
                    ${dataBlog[index].content}
                    </p>
                    <p style="font-size: 15px; margin-left: 230px;">
                                         ${getDistanceTime(dataBlog[index].postAt)}
                    </p>
                    <a href="https://www.instagram.com/">
                        <img
                        style="
                        width: 5%; image-"
                        src="Image/IG.png"
                        />
                    </a>
                    <a href="https://www.facebook.com/">
                        <img
                    style="
                    width: 5%;"
                    src="Image/FB.png"
                    />
                    </a>
                    <a href="https://www.twitter.com/">
                        <img
                        style="
                        width: 5%;"
                        src="Image/TWT.png"
                        />
                    </a>
                    <a href="https://www.twitter.com/">
                        <img
                        style="
                        width: 5%;"
                        src="Image/YT.png"
                        />
                    </a>
                    <div class="d-column">
                        <button class="bg-danger px-4">Edit Post</button>
                        <buttom class="bg-gray px-4">Delete Post</buttom>
                    </div>
                </div>
            </div>
        </div>
    </div>
    `;
  }
}

function getFullTime(time) {
  // console.log("get full time");
  // let time = new Date();
  // console.log(time);

  let monthName = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];

  let date = time.getDate();

  let monthIndex = time.getMonth();

  let year = time.getFullYear();

  let hours = time.getHours();
  
  let minutes = time.getMinutes();

  if (hours <= 9) {
    hours = "0" + hours;
  } else if (minutes <= 9) {
    minutes = "0" + minutes;
  }

  return `${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WIB`;
}

function getDistanceTime(time) {
  let timeNow = new Date();
  let timePost = time;

  let distance = timeNow - timePost;
  console.log(distance);

  let milisecond = 1000;
  let secondInHours = 3600;
  let hoursInDays = 24;

  let distanceDay = Math.floor(
    distance / (milisecond * secondInHours * hoursInDays));
  let distanceHours = Math.floor(distance / (milisecond * 60 * 60));
  let distanceMinutes = Math.floor(distance / (milisecond * 60));
  let distanceSeconds = Math.floor(distance / milisecond);

  if (distanceDay > 0) {
    return `${distanceDay} Days Ago`;
  } else if (distanceHours > 0) {
    return `${distanceHours} Hours Ago`;
  } else if (distanceMinutes > 0) {
    return `${distanceMinutes} Minutes Ago`;
  } else {
    return `${distanceSeconds} Seconds Ago`;
  }
}

setInterval(function () {
  renderBlog();
}, 10000);