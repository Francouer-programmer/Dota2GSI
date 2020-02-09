var camera, scene, renderer,
    geometry, material, mesh;

var CANVAS_WIDTH = 300;
var CANVAS_HEIGHT = 300;


initClassElements();

function pairArraysIntoDict(names, classes) {
    
    
}

function initClassElements() {
    var array_1 = new Array("fee","fie","foo");
    var Values = Object.values(document.getElementsByClassName("two"))
    console.log(Values)
    for (i in Values) {
        array_1.forEach(name => init(Values[i], name));
        animate(name);
    }
}

function init(element, name) {
    stats = new Stats();
    stats.setMode(0);
    stats.domElement.style.position = 'absolute';
    stats.domElement.style.left = '0px';
    stats.domElement.style.top = '0px';

    element.appendChild(stats.domElement);

    clock = new THREE.Clock();

    renderer = new THREE.WebGLRenderer();
    renderer.setClearColor(0xffffff, 0.5);
    renderer.setSize(window.innerWidth, window.innerHeight);

    scene = new THREE.Scene();
    

    camera = new THREE.PerspectiveCamera(75, CANVAS_WIDTH / CANVAS_HEIGHT, 1, 10000);
    camera.position.z = 1000;
    scene.add(camera);

    // RENDERER

    renderer = new THREE.WebGLRenderer();
    renderer.setClearColor(0x000, 1.0);
    renderer.setSize(CANVAS_WIDTH, CANVAS_HEIGHT);

    geometry = new THREE.BoxGeometry(200, 200, 200);
    material = new THREE.MeshLambertMaterial({ color: 0xa00ff, wireframe: false });
    mesh = new THREE.Mesh(geometry, material);
    //scene.add( mesh );
    cubeSineDriver = 0;

    textGeo = new THREE.PlaneGeometry(300, 300);
    THREE.ImageUtils.crossOrigin = ''; //Need this to pull in crossdomain images from AWS
    textTexture = THREE.ImageUtils.loadTexture('http://localhost:3000/assets/img/quickText.png');
    textMaterial = new THREE.MeshLambertMaterial({ color: 0x00ffff, opacity: 0.5, map: textTexture, transparent: true, blending: THREE.AdditiveBlending })
    text = new THREE.Mesh(textGeo, textMaterial);
    text.position.z = 800;
    scene.add(text);

    light = new THREE.DirectionalLight(0xffffff, 0.5);
    light.position.set(-1, 0, 1);
    scene.add(light);

    smokeTexture = THREE.ImageUtils.loadTexture('http://localhost:3000/assets/img/Smoke-Element.png');
    smokeMaterial = new THREE.MeshLambertMaterial({ color: 0x00dddd, map: smokeTexture, transparent: true });
    smokeGeo = new THREE.PlaneGeometry(300, 300);
    smokeParticles = [];


    for (p = 0; p < 100; p++) {
        var particle = new THREE.Mesh(smokeGeo, smokeMaterial);
        particle.position.set(Math.random() * 500 - 250, Math.random() * 500 - 250, Math.random() * 1000 - 100);
        particle.rotation.z = Math.random() * 360;
        scene.add(particle);
        smokeParticles.push(particle);
    }

    element.appendChild(renderer.domElement);

}

function animate() {

    // note: three.js includes requestAnimationFrame shim
    stats.begin();
    delta = clock.getDelta();
    requestAnimationFrame(animate);
    evolveSmoke();
    render();
    stats.end();
}

function evolveSmoke() {
    var sp = smokeParticles.length;
    while (sp--) {
        smokeParticles[sp].rotation.z += (delta * 0.2);
    }
}

function render() {

    mesh.rotation.x += 0.005;
    mesh.rotation.y += 0.01;
    cubeSineDriver += .01;
    mesh.position.z = 100 + (Math.sin(cubeSineDriver) * 500);
    renderer.render(scene, camera);

}